import BN from "bn.js";
import { Account, TransactionConfig, TransactionReceipt } from "web3-core";
import { ContractAddresses } from "../../deployment/tasks/common";

import Web3 from "web3";
import { FTSOParameters } from "../../deployment/config/FTSOParameters";
import { VoterRegistry, FlareSystemManager } from "../../typechain-web3/contracts/fastUpdates/interface/mocks";
import { FlareSystemMock } from "../../typechain-web3/contracts/fastUpdates/test";
import { FastUpdater } from "../../typechain-web3/contracts/fastUpdates/implementation/FastUpdater";
import { FastUpdateIncentiveManager } from "../../typechain-web3/contracts/fastUpdates/implementation/FastUpdateIncentiveManager";
import { NonPayableTransactionObject } from "../../typechain-web3/types";
import { toBN } from "../utils/voting-utils";
import { getAccount, loadContract } from "../utils/web3";
import { getLogger } from "../utils/logger";
import { promiseWithTimeout, retryPredicate } from "../utils/retry";
import { RevertedTxError, asError } from "../utils/error";
import { SortitionKey, Proof } from "../Sortition";
import { sleepFor } from "../utils/time";
import { BareSignature } from "../utils/voting-types";

interface TypeChainContracts {
    readonly voterRegistry: VoterRegistry;
    readonly flareSystemManager: FlareSystemManager;
    readonly fastUpdater: FastUpdater;
    readonly fastUpdateIncentiveManager: FastUpdateIncentiveManager;
    readonly flareSystemMock: FlareSystemMock;
}

export class Web3Provider {
    private readonly logger = getLogger(Web3Provider.name);
    public readonly account: Account;

    private constructor(
        readonly contractAddresses: ContractAddresses,
        readonly web3: Web3,
        private contracts: TypeChainContracts,
        private config: FTSOParameters,
        privateKey: string
    ) {
        this.account = getAccount(this.web3, privateKey);
    }

    // only on mocked contract
    async registerAsAVoter(epoch: number, key: SortitionKey, weight: number, address: string, addToNonce?: number) {
        const x = "0x" + "0".repeat(64 - key.pk.x.toString(16).length) + key.pk.x.toString(16);
        const y = "0x" + "0".repeat(64 - key.pk.y.toString(16).length) + key.pk.y.toString(16);

        const newProvider: [string, string, number] = [x, y, weight];
        const methodCall = this.contracts.flareSystemMock.methods.registerAsVoter(epoch, address, newProvider);
        return this.signAndFinalize(
            "RegisterAsAVoter",
            this.contracts.voterRegistry.options.address,
            methodCall,
            undefined,
            undefined,
            addToNonce
        );
    }

    // only by daemon
    async advanceIncentive(addToNonce?: number) {
        const methodCall = this.contracts.fastUpdateIncentiveManager.methods.advance();
        return this.signAndFinalize(
            "advance",
            this.contracts.fastUpdateIncentiveManager.options.address,
            methodCall,
            undefined,
            undefined,
            addToNonce
        );
    }

    // only by daemon
    async freeSubmitted(addToNonce?: number) {
        const methodCall = this.contracts.fastUpdater.methods.freeSubmitted();
        return this.signAndFinalize(
            "freeSubmitted",
            this.contracts.fastUpdater.options.address,
            methodCall,
            undefined,
            undefined,
            addToNonce
        );
    }

    async getWeight(address: string) {
        const weight = await this.contracts.fastUpdater.methods.currentSortitionWeight(address).call();

        return weight;
    }

    async fetchCurrentPrices(feeds: number[]) {
        // console.log("address", address);
        const prices = await this.contracts.fastUpdater.methods.fetchCurrentPrices(feeds).call();

        return prices;
    }

    async getBaseSeed(): Promise<BN> {
        const seed = await this.contracts.flareSystemManager.methods.getCurrentRandom().call();

        return toBN(seed);
    }
    async getCurrentScoreCutoff(): Promise<string> {
        const cutoff = await this.contracts.fastUpdater.methods.currentScoreCutoff().call();

        return cutoff;
    }
    async submitUpdates(
        proof: Proof,
        replicate: number,
        deltas: [string[], string],
        submissionBlockNum: number,
        signature: BareSignature,
        addToNonce?: number
    ) {
        const sortitionCredential: [number, [string, string], string, string] = [
            replicate,
            [proof.gamma.x.toString(), proof.gamma.y.toString()],
            proof.c.toString(),
            proof.s.toString(),
        ];

        const newFastUpdate: [
            number,
            [number, [string, string], string, string],
            [string[], string],
            [number, string, string]
        ] = [submissionBlockNum, sortitionCredential, deltas, [signature.v, signature.r, signature.s]];

        const methodCall = this.contracts.fastUpdater.methods.submitUpdates(newFastUpdate);

        return await this.signAndFinalize(
            "submitUpdates",
            this.contracts.fastUpdater.options.address,
            methodCall,
            undefined,
            undefined,
            addToNonce
        );
    }

    async getBlockNumber(): Promise<number> {
        return this.web3.eth.getBlockNumber();
    }

    public async getNonce(account: Account): Promise<number> {
        return await this.web3.eth.getTransactionCount(account.address);
    }

    private async signAndFinalize(
        label: string,
        toAddress: string,
        fnToEncode: NonPayableTransactionObject<void>,
        value: number | BN = 0,
        from: Account = this.account,
        addToNonce: number = 0,
        gasPriceMultiplier: number = this.config.gasPriceMultiplier
    ): Promise<TransactionReceipt> {
        // Try a dry-run of the transaction first.
        // If it fails, we just return the error and skip sending a real transaction.
        // If it succeeds, we try to send the real transaction. Note that it might still fail if the state changes in the meantime.
        const dryRunError: Error | undefined = await this.dryRunTx(fnToEncode, from, label);
        if (dryRunError instanceof Error) {
            throw dryRunError;
        }

        const gasPrice = toBN(await this.web3.eth.getGasPrice());

        let txNonce: number;
        // console.log("gas price", gasPrice.muln(gasPriceMultiplier).toString(), gasPriceMultiplier);
        const sendTx = async () => {
            txNonce = (await this.getNonce(from)) + addToNonce;
            // console.log("nonce", txNonce);
            const tx = <TransactionConfig>{
                from: from.address,
                to: toAddress,
                gas: this.config.gasLimit.toString(),
                gasPrice: gasPrice.muln(gasPriceMultiplier).toString(),
                data: fnToEncode.encodeABI(),
                value: value,
                nonce: txNonce,
            };
            const signedTx = await from.signTransaction(tx);
            return this.web3.eth.sendSignedTransaction(signedTx.rawTransaction!);
        };

        const receipt: TransactionReceipt = await promiseWithTimeout(sendTx(), 20000);

        const isTxFinalized = async () => (await this.getNonce(from)) > txNonce;
        await retryPredicate(isTxFinalized, 8, 1000);

        return receipt;
    }

    /**
     * Web3js can be configured to include the revert reason in the original error object when sending a transaction (using `web3.eth.handleRevert`).
     * However, this doesn't seem to work reliably - it does not always produce the revert reason. So instead, we run the dry-run logic again,
     * where it seems to be always populated).
     *
     * Note that this still might fail if the state changes between the original transaction call and the dry-run call, but that is not very likely to happen.
     */
    private async getRevertReasonError(
        label: string,
        fnToEncode: NonPayableTransactionObject<void>,
        from: Account
    ): Promise<Error> {
        const revertError = await this.dryRunTx(fnToEncode, from, label);
        if (revertError === undefined) {
            return new RevertedTxError(
                `[${label}] Transaction reverted, but failed to get revert reason - state might have changed since original transaction call.`
            );
        } else return revertError;
    }

    /** Simulates running the transaction. Returns an error object with the revert reason if it fails. */
    private async dryRunTx(
        fnToEncode: NonPayableTransactionObject<void>,
        from: Account,
        label: string
    ): Promise<Error | undefined> {
        try {
            await fnToEncode.call({ from: from.address });
            return undefined;
        } catch (e: unknown) {
            return new RevertedTxError(`[${label}] Transaction reverted`, asError(e));
        }
    }

    static async create(contractAddresses: ContractAddresses, web3: Web3, config: FTSOParameters, privateKey: string) {
        const contracts = {
            voterRegistry: await loadContract<VoterRegistry>(web3, contractAddresses.voterRegistry, "VoterRegistry"),
            flareSystemManager: await loadContract<FlareSystemManager>(
                web3,
                contractAddresses.flareSystemManager,
                "FlareSystemManager"
            ),
            fastUpdater: await loadContract<FastUpdater>(web3, contractAddresses.fastUpdater, "FastUpdater"),
            flareSystemMock: await loadContract<FlareSystemMock>(
                web3,
                contractAddresses.flareSystemManager,
                "FlareSystemMock"
            ),
            fastUpdateIncentiveManager: await loadContract<FastUpdateIncentiveManager>(
                web3,
                contractAddresses.fastUpdateIncentiveManager,
                "FastUpdateIncentiveManager"
            ),
        };

        const provider = new Web3Provider(contractAddresses, web3, contracts, config, privateKey);
        return provider;
    }

    async waitForNewEpoch(epochLen: number) {
        const blockNum = await this.getBlockNumber();
        const newEpochBlockNum = (Math.floor(blockNum / epochLen) + 1) * epochLen;

        await this.waitForBlock(newEpochBlockNum);
    }

    async waitForBlock(minedBlockNum: number) {
        for (;;) {
            const blockNum = await this.getBlockNumber();
            if (blockNum >= minedBlockNum) {
                return;
            }
            await sleepFor(1000);
        }
    }
}

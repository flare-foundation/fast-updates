import { getLogger } from "./utils/logger";
import { Web3Provider } from "./providers/Web3Provider";
import { KeyGen } from "./Sortition";
import { VerifiableRandomness, Randomness, Proof } from "./Sortition";
import { PriceFeedProvider } from "./price-feeds/PriceFeedProvider";
import { encodePacked } from "web3-utils";
import { sha256 } from "ethers";
import { signMessage } from "./utils/web3";

export class PriceVoter {
    private readonly logger = getLogger(PriceVoter.name);
    private readonly key = KeyGen();
    private address: string;
    private epochLen: number;
    private weight: number;
    private priceFeedProvider: PriceFeedProvider;
    private lastRegisteredEpoch: number;
    private privateKey: string;

    constructor(
        private readonly provider: Web3Provider,
        priceFeedProvider: PriceFeedProvider,
        address: string,
        epochLen: number,
        weight: number,
        privateKey: string
    ) {
        this.provider = provider;
        this.priceFeedProvider = priceFeedProvider;
        this.address = address;
        this.epochLen = epochLen;
        this.weight = weight;
        this.lastRegisteredEpoch = -1;
        this.privateKey = privateKey;
    }

    async registerAsVoter(epoch: number, weight: number, addToNonce?: number) {
        const receipt = await this.provider.registerAsAVoter(epoch, this.key, weight, this.address, addToNonce);
        this.lastRegisteredEpoch = epoch;
        return receipt;
    }

    async submitUpdates(
        proof: Proof,
        replicate: number,
        deltas: [string[], string],
        submissionBlockNum: number,
        addToNonce?: number
    ) {
        const toHash = encodePacked(
            { value: submissionBlockNum.toString(), type: "uint256" },
            { value: replicate.toString(), type: "uint256" },
            { value: proof.gamma.x.toString(), type: "uint256" },
            { value: proof.gamma.y.toString(), type: "uint256" },
            { value: proof.c.toString(), type: "uint256" },
            { value: proof.s.toString(), type: "uint256" },
            { value: deltas[0][0], type: "bytes32" },
            { value: deltas[0][1], type: "bytes32" },
            { value: deltas[0][2], type: "bytes32" },
            { value: deltas[0][3], type: "bytes32" },
            { value: deltas[0][4], type: "bytes32" },
            { value: deltas[0][5], type: "bytes32" },
            { value: deltas[0][6], type: "bytes32" },
            { value: deltas[1], type: "bytes32" }
        )!;
        const signature = signMessage(this.provider.web3, sha256(toHash), this.privateKey);

        const receipt = await this.provider.submitUpdates(
            proof,
            replicate,
            deltas,
            submissionBlockNum,
            signature,
            addToNonce
        );
        return receipt;
    }

    async getWeight(address: string) {
        return await this.provider.getWeight(address);
    }

    async getMyWeight() {
        return await this.provider.getWeight(this.address);
    }

    async fetchCurrentPrices(feeds: number[]) {
        return await this.provider.fetchCurrentPrices(feeds);
    }

    async run() {
        let myWeight: number = 0;
        let currentRandom: bigint = BigInt(0);

        for (;;) {
            const blockNum = await this.provider.getBlockNumber();
            const epoch = Math.floor((blockNum + 1) / this.epochLen);

            // just before the epoch ends, register for the new epoch
            if ((blockNum + 1) % this.epochLen >= this.epochLen - 4 && epoch + 1 > this.lastRegisteredEpoch) {
                await this.reRegister(epoch + 1);
                await this.provider.waitForBlock(blockNum + 1);
                continue;
            }

            // client cannot participate yet, since it is not registered for this epoch
            if (epoch > this.lastRegisteredEpoch) {
                this.logger.info(`waiting for epoch ${epoch + 1}, current block ${blockNum}`);
                await this.provider.waitForBlock(blockNum - (blockNum % this.epochLen) + this.epochLen - 4);
                continue;
            }

            // to late to submit, epoch will change in the next block
            if (blockNum % this.epochLen == this.epochLen - 1) {
                continue;
            }

            // new epoch started, get client's weight and the new random seed
            if (blockNum % this.epochLen == 0) {
                myWeight = Number(await this.getMyWeight());
                currentRandom = BigInt((await this.provider.getBaseSeed()).toString());
                this.logger.info(
                    `new epoch ${epoch}, my weight in this epoch: ${myWeight}, current random: ${currentRandom}`
                );
            }

            await this.tryToSubmitUpdates(myWeight, blockNum, currentRandom);

            await this.provider.waitForBlock(blockNum + 1);
        }
    }

    async tryToSubmitUpdates(myWeight: number, blockNum: number, seed: bigint) {
        let addToNonce = 0;
        const cutoff = await this.provider.getCurrentScoreCutoff();

        for (let rep = 0; rep < myWeight; rep++) {
            const r: bigint = Randomness(this.key, seed, BigInt(blockNum), BigInt(rep));

            const newBlockNum = await this.provider.getBlockNumber();
            if (newBlockNum > blockNum) {
                // console.log("time is up for block", blockNum, "got to rep", rep);
                break;
            }

            if (r < BigInt(cutoff)) {
                const proof: Proof = VerifiableRandomness(this.key, seed, BigInt(blockNum), BigInt(rep));

                const deltas: [[string[], string], string] = this.priceFeedProvider.getFeed();
                this.logger.info(`submitting update ${deltas[1]} with rep ${rep} for block ${blockNum}`);
                this.submitUpdates(proof, rep, deltas[0], blockNum, addToNonce)
                    .then(receipt => {
                        const status = receipt.status ? "successful" : "fail";

                        this.logger.info(`update ${status} in block ${receipt.blockNumber}`);
                    })
                    .catch(error => {
                        this.logger.error(`failed to submit updates ${error}`);
                    });
                addToNonce++;
            }
        }
    }

    async reRegister(epoch: number) {
        let status = true;

        // register for the next epoch
        const promise = this.registerAsVoter(epoch, this.weight)
            .then(receipt => {
                this.logger.info(
                    `registered for epoch ${epoch} in block ${receipt.blockNumber} status: ${receipt.status}`
                );
            })
            .catch(error => {
                this.logger.error(`failed to register as a voter for the new epoch ${error}`);
                status = false;
            });

        await promise;

        return status;
    }
}

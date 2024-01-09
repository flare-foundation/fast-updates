import { getLogger } from "./utils/logger";
import { sleepFor } from "./utils/time";
import { Web3Provider } from "./providers/Web3Provider";
import { KeyGen } from "./Sortition";
import { VerifiableRandomness, Proof } from "./Sortition";
import { TransactionReceipt } from "web3-core";
import { PriceFeedProvider } from "./price-feeds/PriceFeedProvider";

export class PriceVoter {
    private readonly logger = getLogger(PriceVoter.name);
    private readonly key = KeyGen();
    private address: string;
    private epochLen: number;
    private weight: number;
    private priceFeedProvider: PriceFeedProvider;

    constructor(
        private readonly provider: Web3Provider,
        priceFeedProvider: PriceFeedProvider,
        address: string,
        epochLen: number,
        weight: number
    ) {
        this.provider = provider;
        this.priceFeedProvider = priceFeedProvider;
        this.address = address;
        this.epochLen = epochLen;
        this.weight = weight;
    }

    async registerAsVoter(epoch: number, weight: number, forceNonce?: number) {
        return await this.provider.registerAsAVoter(epoch, weight, forceNonce);
    }

    async registerAsProvider(forceNonce?: number) {
        const baseSeed = await this.provider.getBaseSeed();
        const replicate = BigInt(0); // Registration doesn't allow cherry-picking a replicate
        const proof: Proof = VerifiableRandomness(this.key, BigInt(baseSeed.toString()), replicate);

        return await this.provider.registerAsProvider(this.key, proof, forceNonce);
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

    async run(feeds: number[]) {
        console.log("waiting for new epoch");
        await this.waitForNewEpoch();

        let myWeight: number = 0;
        for (;;) {
            let addToNonce = 0;
            const blockNum = await this.provider.getBlockNumber();

            let voterReceipt: Promise<TransactionReceipt>;
            let providerReceipt: Promise<TransactionReceipt>;
            const updateReceipts: Promise<TransactionReceipt>[] = [];
            let receipt: TransactionReceipt = <TransactionReceipt>{};
            if ((blockNum + 1) % this.epochLen == 0) {
                console.log("new epoch", (blockNum + 1) / this.epochLen);
                myWeight = Number(await this.getMyWeight());
                console.log("my weight in this epoch weight", myWeight);
            }

            const currentPrices = await this.fetchCurrentPrices(feeds);
            console.log("current prices", currentPrices);

            const sortitionRound = await this.provider.getSortitionRound(blockNum);

            for (let rep = 0; rep < myWeight; rep++) {
                const proof: Proof = VerifiableRandomness(this.key, BigInt(sortitionRound.seed), BigInt(rep));

                if (proof.gamma.x < BigInt(sortitionRound.cutoff)) {
                    const deltas: [string[], string] = this.priceFeedProvider.getFeed();
                    console.log(
                        "submitting",
                        deltas[0][0].slice(2, 2 + Math.ceil(this.priceFeedProvider.numFeeds / 2)),
                        rep,
                        "for block",
                        blockNum
                    );

                    updateReceipts.push(this.provider.submitUpdates(proof, rep, deltas, blockNum, addToNonce));
                    addToNonce++;
                }
                if (updateReceipts.length >= 2) {
                    break;
                }
            }

            if ((blockNum + 1) % this.epochLen == 0) {
                // register for the next epoch
                const nextEpoch = (blockNum + 1) / this.epochLen + 1;
                voterReceipt = this.registerAsVoter(nextEpoch, this.weight, addToNonce);
                addToNonce++;

                providerReceipt = this.registerAsProvider(addToNonce);
                addToNonce++;

                receipt = await voterReceipt;
                console.log(
                    "registered again, now for epoch",
                    nextEpoch,
                    "in block",
                    receipt.blockNumber,
                    "status",
                    receipt.status
                );
                receipt = await providerReceipt;
                console.log(
                    "registered as a provider for the next epoch",
                    "in block",
                    receipt.blockNumber,
                    "status",
                    receipt.status
                );
            }

            for (let i = 0; i < updateReceipts.length; i++) {
                const updateReceipt = updateReceipts[i];
                const receipt = await updateReceipt;
                console.log("update", "in block", receipt.blockNumber, "status", receipt.status);
            }

            await this.waitForBlock(blockNum + 1);
        }
    }

    async waitForNewEpoch() {
        const blockNum = await this.provider.getBlockNumber();
        const newEpochBlockNum = (Math.floor(blockNum / this.epochLen) + 1) * this.epochLen;

        await this.waitForBlock(newEpochBlockNum - 1);
    }

    async waitForBlock(minedBlockNum: number) {
        for (;;) {
            const blockNum = await this.provider.getBlockNumber();
            if (blockNum >= minedBlockNum) {
                return;
            }
            await sleepFor(1000);
        }
    }
}

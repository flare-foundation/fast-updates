import { FTSOClient } from "./FTSOClient";
import { getLogger } from "./utils/logger";
import { sleepFor } from "./utils/time";
import { errorString } from "./utils/error";
import { BlockIndexer } from "./BlockIndexer";
import { EpochSettings } from "./protocol/utils/EpochSettings";
import { EpochData } from "./protocol/voting-types";
import { Web3Provider } from "./providers/Web3Provider";
import { KeyGen } from "./Sortition";
import { VerifiableRandomness, SortitionKey, Proof } from "./Sortition";
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
            // let updateReceipt: Promise<TransactionReceipt | void> = Promise.resolve();
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
            // console.log("sortition round", sortitionRound);

            for (let rep = 0; rep < myWeight; rep++) {
                const proof: Proof = VerifiableRandomness(this.key, BigInt(sortitionRound.seed), BigInt(rep));

                if (proof.gamma.x < BigInt(sortitionRound.cutoff)) {
                    // console.log("submitting +-0-0+ with rep", rep, "for block", blockNum);
                    // const delta1 = "0x7310000000000000000000000000000000000000000000000000000000000000";
                    const delta3 = "0x0000000000000000000000000000000000000000000000000000000000000000";
                    const delta2 = "0x0000000000000000000000000000000000000000000000000000";
                    // const deltas: [string[], string] = [
                    //     [delta1, delta1, delta1, delta1, delta1, delta1, delta1],
                    //     delta2,
                    // ];
                    const ddeltas = this.priceFeedProvider.getFeed();
                    const deltas: [string[], string] = [
                        ["0x" + ddeltas.slice(0, 64), delta3, delta3, delta3, delta3, delta3, delta3],
                        delta2,
                    ];
                    console.log("submitting", ddeltas.slice(0, 3), rep, "for block", blockNum);

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
            await sleep(1000);
        }
    }

    scheduleActions() {
        // const timeSec = this.currentTimeSec();
        // const nextEpochStartSec = this.epochs.nextPriceEpochStartSec(timeSec);

        setTimeout(async () => {
            const myWeight = await this.getMyWeight();
            console.log("my weight", myWeight);
            this.scheduleActions();
            // try {
            //     await this.onPriceEpoch(); // TODO: If this runs for a long time, it might get interleaved with the next price epoch - is this a problem?
            // } catch (e) {
            //     this.logger.error(`Error in price epoch, terminating: ${errorString(e)}`);
            //     process.exit(1);
            // }
        }, 5000);
    }

    // async onPriceEpoch() {
    //     const currentPriceEpochId = this.epochs.priceEpochIdForTime(this.currentTimeSec());

    //     if (
    //         this.lastProcessedPriceEpochId !== undefined &&
    //         this.lastProcessedPriceEpochId !== currentPriceEpochId - 1
    //     ) {
    //         this.logger.error(
    //             `Skipped a price epoch. Last processed: ${this.lastProcessedPriceEpochId}, current: ${currentPriceEpochId}. Will to participate in this round.`
    //         );
    //         this.previousPriceEpochData = undefined;
    //     }

    //     const currentRewardEpochId = this.epochs.rewardEpochIdForPriceEpochId(currentPriceEpochId);
    //     this.logger.info(
    //         `[${currentPriceEpochId}] Processing price epoch, current reward epoch: ${currentRewardEpochId}.`
    //     );

    //     const nextRewardEpochId = currentRewardEpochId + 1;

    //     if (this.isRegisteredForRewardEpoch(currentRewardEpochId)) {
    //         await this.runVotingProcotol(currentPriceEpochId);
    //         this.lastProcessedPriceEpochId = currentPriceEpochId;
    //     }

    //     await this.maybeRegisterForRewardEpoch(nextRewardEpochId);

    //     this.logger.info(`[${currentPriceEpochId}] Finished processing price epoch.`);
    // }

    // private async runVotingProcotol(currentEpochId: number) {
    //     const priceEpochData = this.client.getPricesForEpoch(currentEpochId);
    //     this.logger.info(`[${currentEpochId}] Committing data for current epoch.`);
    //     await this.client.commit(priceEpochData);

    //     await sleepFor(2000);
    //     if (this.previousPriceEpochData !== undefined) {
    //         const previousEpochId = currentEpochId - 1;
    //         this.logger.info(`[${currentEpochId}] Revealing data for previous epoch: ${previousEpochId}.`);
    //         await this.client.reveal(this.previousPriceEpochData);
    //         await this.waitForRevealEpochEnd();
    //         this.logger.info(
    //             `[${currentEpochId}] Calculating results for previous epoch ${previousEpochId} and signing.`
    //         );
    //         const result = await this.client.calculateResultsAndSign(previousEpochId);
    //         await this.awaitFinalization(previousEpochId);

    //         await this.client.publishPrices(result, [0, 1]);
    //     }
    //     this.previousPriceEpochData = priceEpochData;
    // }

    // private async awaitFinalization(priceEpochId: number) {
    //     while (!this.index.getFinalize(priceEpochId)) {
    //         this.logger.info(`Epoch ${priceEpochId} not finalized, keep processing new blocks`);
    //         await sleepFor(500);
    //     }
    //     this.logger.info(`Epoch ${priceEpochId} finalized, continue.`);
    // }

    // private async maybeRegisterForRewardEpoch(nextRewardEpochId: number) {
    //     if (
    //         this.isRegisteredForRewardEpoch(nextRewardEpochId) ||
    //         this.index.getRewardOffers(nextRewardEpochId).length === 0
    //     ) {
    //         return;
    //     }
    //     this.logger.info(`Registering for next reward epoch ${nextRewardEpochId}`);
    //     await this.client.registerAsVoter(nextRewardEpochId);

    //     this.registeredRewardEpochs.add(nextRewardEpochId);
    // }

    // private isRegisteredForRewardEpoch(rewardEpochId: number): boolean {
    //     return this.registeredRewardEpochs.has(rewardEpochId);
    // }

    // private async waitForRevealEpochEnd() {
    //     const revealPeriodDurationMs = this.epochs.revealDurationSec * 1000;
    //     await sleepFor(revealPeriodDurationMs + 1);
    // }

    // private currentTimeSec(): number {
    //     return Math.floor(Date.now() / 1000);
    // }
}

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

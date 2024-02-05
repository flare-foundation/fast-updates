import { getLogger } from "./utils/logger";
import { Web3Provider } from "./providers/Web3Provider";
import { KeyGen } from "./Sortition";
import { VerifiableRandomness, Randomness, Proof } from "./Sortition";
import { PriceFeedProvider } from "./price-feeds/PriceFeedProvider";

export class PriceVoter {
    private readonly logger = getLogger(PriceVoter.name);
    private readonly key = KeyGen();
    private address: string;
    private epochLen: number;
    private weight: number;
    private priceFeedProvider: PriceFeedProvider;
    private lastRegisteredEpoch: number;

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
        this.lastRegisteredEpoch = -1;
    }

    async registerAsVoter(epoch: number, weight: number, addToNonce?: number) {
        const receipt = await this.provider.registerAsAVoter(epoch, this.key, weight, addToNonce);
        this.lastRegisteredEpoch = epoch;
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
        this.logger.info("waiting for a new epoch");
        await this.provider.waitForNewEpoch(this.epochLen);

        let myWeight: number = 0;
        let currentRandom: bigint = BigInt(0);

        for (;;) {
            const blockNum = await this.provider.getBlockNumber();
            const epoch = Math.floor((blockNum + 1) / this.epochLen);

            if (blockNum % this.epochLen >= this.epochLen - 4) {
                if (epoch + 1 > this.lastRegisteredEpoch) {
                    const status = await this.reRegister(epoch);
                    if (status) {
                        await this.provider.waitForNewEpoch(this.epochLen);
                    } else {
                        await this.provider.waitForBlock(blockNum + 1);
                    }
                    continue;
                }
            }

            if (blockNum % this.epochLen == 0) {
                myWeight = Number(await this.getMyWeight());
                currentRandom = BigInt((await this.provider.getBaseSeed()).toString());
                this.logger.info(
                    `new epoch ${epoch}, my weight in this epoch weight: ${myWeight}, current random: ${currentRandom}`
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
                this.provider
                    .submitUpdates(proof, rep, deltas[0], blockNum, addToNonce)
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
        // console.log("end of epoch procedure");
        let status = true;

        // register for the next epoch
        const promise = this.registerAsVoter(epoch + 1, this.weight)
            .then(receipt => {
                this.logger.info(
                    `registered again, now for epoch ${epoch + 1} in block ${receipt.blockNumber} status: ${
                        receipt.status
                    }`
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

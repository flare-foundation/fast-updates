import { readFileSync } from 'fs'

import { feedPaths } from '../config'
import { getOrCreateLogger } from '../utils'
import type { PriceDeltas } from '../utils/'

type PriceFeed = {
    [timestampSec: string]: number
}

export class ExamplePriceFeedProvider {
    private readonly logger = getOrCreateLogger(ExamplePriceFeedProvider.name)
    private readonly priceFeeds: PriceFeed[] = []
    // Rounding time for price feed timestamps in seconds
    private readonly priceTsRounding = 5

    constructor(
        private readonly numFeeds: number,
        private readonly precision: number = 0,
        private readonly genesisTimeSec: number = 0
    ) {
        if (numFeeds > 1) {
            throw new Error(
                'Number of feeds should be at most 1. Temporary limit in place.'
            )
        }
        if (precision === 0) {
            this.logger.info(
                `PriceFeedProvider (observer) initialized with ${numFeeds} feeds at genesis time ${genesisTimeSec}`
            )
        } else {
            this.logger.info(
                `PriceFeedProvider initialized with ${numFeeds} feeds and precision ${precision} at genesis time ${genesisTimeSec}`
            )
        }

        for (let i = 0; i < feedPaths.length; i++) {
            const priceFeed = JSON.parse(
                readFileSync(feedPaths[i] as string).toString()
            ) as PriceFeed
            this.priceFeeds.push(priceFeed)
        }
    }

    /**
     * Retrieves the current prices for the given feeds.
     *
     * @param feeds - An array of feed numbers.
     * @returns An array of current prices corresponding to the feeds.
     */
    public getCurrentPrices(feeds: number[]): number[] {
        const currentTimeSec = Math.floor(Date.now() / 1000)
        const diffSec = currentTimeSec - this.genesisTimeSec
        const discreteDiffSec =
            Math.floor(diffSec / this.priceTsRounding) * this.priceTsRounding

        this.logger.debug(
            `currentTime ${currentTimeSec}, genesisTime ${this.genesisTimeSec}, difference ${diffSec}, discretizedDifference ${discreteDiffSec}`
        )

        const currentPrices = []
        for (let i = 0; i < feeds.length; i++) {
            currentPrices.push(
                (this.priceFeeds[i] as PriceFeed)[discreteDiffSec.toString()] ??
                    0
            )
        }
        return currentPrices
    }

    /**
     * Retrieves the feed based on the given chain prices and feed prices.
     * @param onChainPrices - An array of chain prices.
     * @param offChainPrices - An array of feed prices.
     * @returns A tuple containing the deltas and the representation.
     * The deltas consist of an array of strings and a string.
     * The representation is a string.
     * @throws Error if the length of chainPrices and feedPrices is not equal or if the length is not 1.
     */
    public getFastUpdateDeltas(
        onChainPrices: number[],
        offChainPrices: number[]
    ): PriceDeltas {
        if (
            onChainPrices.length != offChainPrices.length ||
            onChainPrices.length != 1
        ) {
            throw new Error('Arrays should be of equal length')
        }

        let feeds: string = ''
        let feed: number = 0
        let rep: string = ''

        for (let i = 0; i < onChainPrices.length; i++) {
            if (i % 2 == 0) {
                feed = 0
            }
            const onChainPrice = onChainPrices[i] as number
            const offChainPrice = offChainPrices[i] as number

            if (onChainPrice > offChainPrice) {
                // If the chain price is greater than the feed price, decrease the chain price
                if (i % 2 == 0) {
                    feed += 12
                    rep += '-'
                } else {
                    feed += 3
                    rep += '-'
                    feeds = feeds + feed.toString(16)
                }
            } else if (onChainPrice < offChainPrice) {
                // If the chain price is less than the feed price, increase the chain price
                if (i % 2 == 0) {
                    feed += 4
                    rep += '+'
                } else {
                    feed += 1
                    rep += '+'
                    feeds = feeds + feed.toString(16)
                }
            }
        }

        if (this.numFeeds % 2 == 1) {
            feeds = feeds + feed.toString(16)
        }

        feeds = feeds + '0'.repeat(64 - feeds.length)
        const delta1 = '0x' + '0'.repeat(64)
        const delta2 = '0x' + '0'.repeat(52)
        const deltas: [string[], string] = [
            ['0x' + feeds, delta1, delta1, delta1, delta1, delta1, delta1],
            delta2,
        ]

        return [deltas, rep]
    }
}

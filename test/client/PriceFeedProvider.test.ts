import { ExamplePriceFeedProvider } from '../../client/providers/ExamplePriceFeedProvider'

describe('PriceFeedProvider', () => {
    describe('getFeed', () => {
        it('should return the price feed and representation for the given chain prices and local prices', () => {
            const provider = new ExamplePriceFeedProvider(1)
            const onChainPrices = [100]
            const offChainPrices = [90]

            const [feed, representation] = provider.getFastUpdateDeltas(
                onChainPrices,
                offChainPrices
            )
            expect(feed).to.equal('0xc0')
            expect(representation).to.equal('-')
        })

        it('should throw an error if the arrays are not of equal length or the length is not 1', () => {
            const provider = new ExamplePriceFeedProvider(1)
            const onChainPrices = [100, 200]
            const offChainPrices = [90]

            expect(() =>
                provider.getFastUpdateDeltas(onChainPrices, offChainPrices)
            ).to.throw('Arrays should be of equal length')
        })
    })
})

import { RangeFPA, SampleFPA } from '../utils'

export const EPOCH_LEN = 20 as const
export const FEEDS = [0] as const
export const WEIGHT = 1000 as const

// Fast Updater params
export const ANCHOR_PRICES = [7560]
export const SUBMISSION_WINDOW = 10 as const

// Incentive Manager params
export const BASE_RANGE = RangeFPA(2 ** -8)
export const BASE_SAMPLE_SIZE = SampleFPA(2)

// Precision = (BASE_RANGE / BASE_SAMPLE_SIZE) * 2^15
// Scale = (1 + BASE_RANGE / BASE_SAMPLE SIZE) * 2^15

export const SAMPLE_INCREASE_LIMIT = SampleFPA(5)
export const RANGE_INCREASE_PRICE = 5 as const
export const DURATION = 8 as const

if (ANCHOR_PRICES.length !== FEEDS.length) {
    throw new Error('Anchor prices and feeds should have same length')
}

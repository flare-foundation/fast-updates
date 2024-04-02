import BN from 'bn.js'

import { RangeOrSampleFPA } from '../utils'

export const EPOCH_LEN = 25
export const WEIGHT = 1000

// Fast Updater params
export const FEEDS = [0, 1, 2, 3, 4, 5] as const
export const ANCHOR_PRICES = [
    new BN(69321.98),
    new BN(3537.78),
    new BN(182.8),
    new BN(19.363),
    new BN(1.6636),
    new BN(0.0000302),
]
export const SUBMISSION_WINDOW = 2

// Incentive Manager params
export const BASE_RANGE = RangeOrSampleFPA(2 ** -8)
export const BASE_SAMPLE_SIZE = RangeOrSampleFPA(2)

// Precision = (BASE_RANGE / BASE_SAMPLE_SIZE) * 2^15
// Scale = (1 + BASE_RANGE / BASE_SAMPLE SIZE) * 2^15

export const SAMPLE_INCREASE_LIMIT = RangeOrSampleFPA(5)
export const RANGE_INCREASE_PRICE = 5
export const DURATION = 8
export const BACKLOG_LEN = 20

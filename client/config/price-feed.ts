import path from 'path'

const BASE_PATH = path.join(path.resolve(), 'client', 'config')

export const feedPaths = [
    path.join(BASE_PATH, 'example-price-feeds', 'offchain_btcusd_prices.json'),
]

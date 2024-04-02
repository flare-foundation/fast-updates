import path from 'path'

const BASE_PATH = path.join(path.resolve(), 'client', 'config')

export const feedPaths = [
    path.join(BASE_PATH, 'example-price-feeds', 'btcusdt.json'),
    path.join(BASE_PATH, 'example-price-feeds', 'ethusdt.json'),
    path.join(BASE_PATH, 'example-price-feeds', 'solusdt.json'),
    path.join(BASE_PATH, 'example-price-feeds', 'linkusdt.json'),
    path.join(BASE_PATH, 'example-price-feeds', 'arbusdt.json'),
    path.join(BASE_PATH, 'example-price-feeds', 'shibusdt.json'),
]

export function RangeFPA(range: number) {
    const r = Math.floor(range * 2 ** 8); // 2^8 since scaled for 2^(-8) for fixed precision arithmetic)
    if (r > 2 ** 16 - 1) {
        throw "range out of bound";
    }
    return r;
}

export function SampleFPA(range: number) {
    const s = Math.floor(range * 2 ** 8); // 2^8 since scaled for 2^(-8) for fixed precision arithmetic)
    if (s > 2 ** 16 - 1) {
        throw "sample out of bound";
    }
    return s;
}

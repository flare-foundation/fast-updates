/**
 * Converts a range value to a fixed-point arithmetic representation.
 * @param range - The range value to be converted.
 * @returns The fixed-point arithmetic representation of the range value.
 * @throws Error if the range value is out of bounds.
 */
export function RangeFPA(range: number): number {
    const r = Math.floor(range * 2 ** 8) // 2^8 since scaled for 2^(-8) for fixed precision arithmetic)
    if (r > 2 ** 16 - 1) {
        throw new Error('range out of bound')
    }
    return r
}

/**
 * Converts a given range to a fixed-point number.
 * @param range The range to be converted.
 * @returns The fixed-point number.
 * @throws Error if the converted number is out of bounds.
 */
export function SampleFPA(range: number): number {
    const s = Math.floor(range * 2 ** 8) // 2^8 since scaled for 2^(-8) for fixed precision arithmetic)
    if (s > 2 ** 16 - 1) {
        throw new Error('sample out of bound')
    }
    return s
}

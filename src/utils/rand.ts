import crypto from "crypto";

export function RandInt(max: bigint) {
    const length = max.toString(2).length;
    const numBytes = Math.floor((length - 1) / 8) + 1;
    const twoToLength = BigInt(2) ** BigInt(length);
    for (;;) {
        const randomBytes = crypto.randomBytes(numBytes).toString("hex");
        const r = BigInt("0x" + randomBytes) % twoToLength;

        if (r < max) {
            return r;
        }
    }
}

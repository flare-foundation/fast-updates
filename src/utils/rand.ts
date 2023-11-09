import crypto from "crypto";

export function RandInt(numBytes: number) {
    const randbytes = crypto.randomBytes(numBytes).toString("hex");
    const r = BigInt("0x" + randbytes);

    return r;
}

import { RandInt } from "../../src/utils/rand";

describe("RandInt", function () {
    it("should generate a random BigInt", async () => {
        for (let i = 0; i < 100; i++) {
            const r = RandInt(BigInt(10));
            expect(r).to.be.lessThan(10);
        }
    });
});

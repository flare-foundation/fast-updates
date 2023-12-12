import chai, { expect } from "chai";
import chaiAsPromised from "chai-as-promised";
import { TestFixedPointArithmeticInstance } from "../../typechain-truffle";
import { getTestFile } from "../../test-utils/utils/constants";
import { RandInt } from "../../src/utils/rand";

chai.use(chaiAsPromised);

const TestFixedPointArithmetic = artifacts.require("TestFixedPointArithmetic");

contract(`FixedPointArithmetic.sol; ${getTestFile(__filename)}`, async accounts => {
    let fpaInstance: TestFixedPointArithmeticInstance;
    before(async () => {
        const governance = accounts[0];
        fpaInstance = await TestFixedPointArithmetic.new(governance);
    });

    // Bit length tests

    it("should have 16 bit Scale values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c1 = await fpaInstance.identityScaleTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identityScaleTest(BigInt(2) ** BigInt(16));
        expect(c2).to.eventually.throw();
    });
    it("should have 16 bit Precision values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c1 = await fpaInstance.identityPrecisionTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identityPrecisionTest(BigInt(2) ** BigInt(16));
        expect(c2).to.eventually.throw();
    });
    it("should have 16 bit SampleSize values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c1 = await fpaInstance.identitySampleSizeTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identitySampleSizeTest(BigInt(2) ** BigInt(16));
        expect(c2).to.eventually.throw();
    });
    it("should have 16 bit Range values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c1 = await fpaInstance.identitySampleSizeTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identitySampleSizeTest(BigInt(2) ** BigInt(16));
        expect(c2).to.eventually.throw();
    });
    it("should have 32 bit Price values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(32) - BigInt(1));
        const c1 = await fpaInstance.identityPriceTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identityPriceTest(BigInt(2) ** BigInt(32));
        expect(c2).to.eventually.throw();
    });
    it("should have signed 8 bit Delta values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(7) - BigInt(1));
        const c1 = await fpaInstance.identityDeltaTest(x);
        const c2 = await fpaInstance.identityDeltaTest(-x);

        expect(x).to.equal(c1);
        expect(-x).to.equal(c2);

        const c3 = fpaInstance.identityDeltaTest(BigInt(2) ** BigInt(7));
        const c4 = fpaInstance.identityDeltaTest(-(BigInt(2) ** BigInt(7) + BigInt(1)));

        expect(c3).to.eventually.throw();
        expect(c4).to.eventually.throw();
    });
    it("should have 16 bit Fractional values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c1 = await fpaInstance.identityFractionalTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identityFractionalTest(BigInt(2) ** BigInt(16));
        expect(c2).to.eventually.throw();
    });
    it("should have 240 bit Fee values", async() => {
        const x = RandInt(BigInt(2) ** BigInt(240) - BigInt(1));
        const c1 = await fpaInstance.identityFeeTest(x);

        expect(x).to.equal(c1);

        const c2 = fpaInstance.identityFeeTest(BigInt(2) ** BigInt(240));
        expect(c2).to.eventually.throw();
    });

    // Arithmetic identity tests
    
    it("should have one as additive one", async () => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c = await fpaInstance.oneTest(x);

        expect(c[0]).to.equal(x);
        expect(c[1]).to.equal(x);
    });
    it("should have zeroD as additive zero", async () => {
        const x = RandInt(BigInt(2) ** BigInt(8) - BigInt(1));
        const c = await fpaInstance.zeroRTest(x);

        expect(c[0]).to.equal(x);
        expect(c[1]).to.equal(x);
    });
    it("should have zeroS as additive zero", async () => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c = await fpaInstance.zeroRTest(x);

        expect(c[0]).to.equal(x);
        expect(c[1]).to.equal(x);
    });
    it("should have zeroR as additive zero", async () => {
        const x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        const c = await fpaInstance.zeroRTest(x);

        expect(c[0]).to.equal(x);
        expect(c[1]).to.equal(x);
    });

    // Addition/subtraction tests

    it("should add and subtract Delta values", async () => {
        const x = RandInt(BigInt(2) ** BigInt(7) - BigInt(1));
        const y = RandInt(BigInt(2) ** BigInt(7) - BigInt(1) - x);

        const c1 = await fpaInstance.addDeltaTest(x, y);
        const c2 = await fpaInstance.addDeltaTest(-x, -y);
        const c3 = await fpaInstance.addDeltaTest(x, -y);

        expect(c1).to.equal(x + y);
        expect(c2).to.equal(-x + (-y));
        expect(c3).to.equal(x + (-y));
    });
    
    it("should add and subtract SampleSize values", async () => {
        var x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        var y = RandInt(BigInt(2) ** BigInt(16) - BigInt(1) - x);
        if (x < y) {
            const z = x;
            x = y;
            y = z;
        }

        const c = await fpaInstance.addSampleSizeTest(x, y);

        expect(c[0]).to.equal(x + y);
        expect(c[1]).to.equal(x - y);
    });

    it("should add and subtract Range values", async () => {
        var x = RandInt(BigInt(2) ** BigInt(16) - BigInt(1));
        var y = RandInt(BigInt(2) ** BigInt(16) - BigInt(1) - x);
        if (x < y) {
            const z = x;
            x = y;
            y = z;
        }

        const c = await fpaInstance.addRangeTest(x, y);

        expect(c[0]).to.equal(x + y);
        expect(c[1]).to.equal(x - y);
    });

    it("should add and subtract Fee values", async () => {
        var x = RandInt(BigInt(2) ** BigInt(240) - BigInt(1));
        var y = RandInt(BigInt(2) ** BigInt(240) - BigInt(1) - x);
        if (x < y) {
            const z = x;
            x = y;
            y = z;
        }

        const c = await fpaInstance.addFeeTest(x, y);

        expect(c[0]).to.equal(x + y);
        expect(c[1]).to.equal(x - y);
    });

    // Multiplication/division tests

    it("should multiply and divide Scale values", async () => {
        const xI = Math.floor(2**15 + (Math.random() * 2**14))
        const yI = Math.floor(2**15 + (Math.random() * 2**14))

        const x = xI / 2**15;
        const y = yI / 2**15;

        const xy1 = Math.floor((x * y) * 2**15) / 2**15
        const xy2 = Math.floor((x / y) * 2**15) / 2**15

        const c = await fpaInstance.mulScaleTest(xI, yI);

        expect(c[0] / 2**15).to.equal(xy1);
        expect(c[1] / 2**15).to.equal(xy2);
    });
    it("should multiply Price and Scale values", async () => {
        const x = Math.floor(Math.random() * 2**31);
        const yI = Math.floor(2**15 + Math.random() * 2**15);
        const y = yI / 2**15;

        const c = await fpaInstance.mulPriceScaleTest(x, yI);

        expect(c).to.equal(Math.floor(x * y));
    });
    it("should multiply Fee and Range values", async () => {
        const x = Math.floor(Math.random() * 2**32);
        const yI = Math.floor(Math.random() * 2**16)
        const y = yI / 2**8;

        const c = await fpaInstance.mulFeeRangeTest(x, yI);

        expect(c).to.equal(Math.floor(x * y));
    })
    it("should multiply Fractional and Fee values", async () => {
        const xI = Math.floor(Math.random() * 2**16)
        const x = xI / 2**16;
        const y = Math.floor(Math.random() * 2**32);

        const c = await fpaInstance.mulFractionalFeeTest(xI, y);

        expect(c).to.equal(Math.floor(x * y));
    })
    it("should multiply Fractional and SampleSize values", async () => {
        const xI = Math.floor(Math.random() * 2**16)
        const x = xI / 2**16;
        const yI = Math.floor(Math.random() * 2**16);
        const y = yI / 2**8;

        const xy = Math.floor(x * y * 2**8) / 2**8

        const c = await fpaInstance.mulFractionalSampleSizeTest(xI, yI);

        expect(c / 2**8).to.equal(xy);
    })
    it("should divide Range values", async () => {
        const yI = Math.floor(Math.random() * 2**16)
        const y = yI / 2**8;
        const xI = Math.floor(Math.random() * yI);
        const x = xI / 2**8;

        const xy = Math.floor((x / y) * 2**16) / 2**16

        const c = await fpaInstance.divRangeTest(xI, yI);

        expect(c / 2**16).to.equal(xy);
    })
    it("should divide Fee values", async () => {
        const y = Math.floor(Math.random() * 2**32)
        const x = Math.floor(Math.random() * y);

        const xy = Math.floor((x / y) * 2**16) / 2**16

        const c = await fpaInstance.divFeeTest(x, y);

        expect(c / 2**16).to.equal(xy);
    })
    it("should divide Range and SampleSize values", async () => {
        const yI = Math.floor(Math.random() * 2**16)
        const y = yI / 2**8;
        const xI = Math.floor(Math.random() * y);
        const x = xI / 2**8;

        const xy = Math.floor((x / y) * 2**15) / 2**15

        const c = await fpaInstance.divRangeSampleSizeTest(xI, yI);

        expect(c / 2**15).to.equal(xy);
    })

});
import {
    RangeOrSampleFPA,
} from '../../deployment/utils/fixed-point-arithmetic'

describe('RangeOrSampleFPA', () => {
    it('should convert a range value to a fixed-point arithmetic representation', () => {
        const range = 0.5
        const result = RangeOrSampleFPA(range)

        expect(result).to.equal("0x8" + "0".repeat(29))
    })

    it('should throw an error if the range value is out of bounds', () => {
        const range = 256

        expect(() => RangeOrSampleFPA(range)).to.throw("range or sample size too large")
    })
})

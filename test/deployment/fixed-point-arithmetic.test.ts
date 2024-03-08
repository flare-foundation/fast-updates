import {
    RangeFPA,
    SampleFPA,
} from '../../deployment/utils/fixed-point-arithmetic'

describe('RangeFPA', () => {
    it('should convert a range value to a fixed-point arithmetic representation', () => {
        const range = 0.5
        const result = RangeFPA(range)

        expect(result).to.equal(128)
    })

    it('should throw an error if the range value is out of bounds', () => {
        const range = 10000

        expect(() => RangeFPA(range)).to.throw('range out of bound')
    })
})

describe('SampleFPA', () => {
    it('should convert a given range to a fixed-point number', () => {
        const range = 0.25
        const result = SampleFPA(range)

        expect(result).to.equal(64)
    })

    it('should throw an error if the converted number is out of bounds', () => {
        const range = 10000

        expect(() => SampleFPA(range)).to.throw('sample out of bound')
    })
})

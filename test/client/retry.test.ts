import { expect } from 'chai'

import { sleepFor } from '../../client/utils/retry'

describe('sleepFor', () => {
    it('should pause execution for the specified number of milliseconds', async () => {
        const ms = 1001
        const start = Date.now()
        await sleepFor(ms)
        const end = Date.now()
        const elapsed = end - start
        expect(elapsed).to.be.greaterThanOrEqual(1000)
    })
})

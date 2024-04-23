/**
 * Represents an error that occurs when a transaction is reverted.
 */
export class TxRevertedError extends Error {
    constructor(message: string, cause?: Error) {
        super(message, { cause: cause })
    }
}

/**
 * Converts an unknown object to an Error.
 * If the object is already an Error, it is returned as is.
 * Otherwise, it throws a new Error with a message indicating that an unknown object was thrown.
 * @param error - The unknown object to convert to an Error.
 * @returns The converted Error object.
 */
export function asError(error: unknown): Error {
    if (error instanceof Error) {
        return error
    } else {
        throw new Error(
            `Unknown object thrown as error: ${JSON.stringify(error)}`
        )
    }
}

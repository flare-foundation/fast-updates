/**
 * Pauses the execution for the specified number of milliseconds.
 * @param ms - The number of milliseconds to sleep for.
 * @returns A promise that resolves after the specified time has elapsed.
 */
export async function sleepFor(ms: number): Promise<void> {
    await new Promise<void>((resolve) => {
        setTimeout(() => {
            resolve()
        }, ms)
    })
}

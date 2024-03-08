import type { TransformableInfo } from 'logform'
import type { Logger } from 'winston'
import winston from 'winston'
import type TransportStream from 'winston-transport'

const logPath = process.env['LOG_PATH'] ? process.env['LOG_PATH'] : './logs/'

const loggers = new Map<string, Logger>()

/**
 * Retrieves an existing logger with the specified label or creates a new one if it doesn't exist.
 * @param label - The label for the logger.
 * @returns The logger instance.
 */
export function getOrCreateLogger(label: string, filename?: string): Logger {
    if (loggers.has(label)) return loggers.get(label) as Logger

    const transports: TransportStream[] = [new winston.transports.Console()]
    if (filename)
        transports.push(
            new winston.transports.File({
                filename: `${logPath}${filename}.log`,
            })
        )

    const logger = winston.createLogger({
        format: winston.format.combine(
            winston.format.timestamp(),
            winston.format.json(),
            winston.format.label({
                label: label,
            }),
            winston.format.printf((json: TransformableInfo) => {
                if (json['label']) {
                    return `${json['timestamp']} - ${json['label']}:[${json.level}]: ${json.message}`
                } else {
                    return `${json['timestamp']} - [${json.level}]: ${json.message}`
                }
            })
        ),
        level: process.env['LOG_LEVEL'] ?? 'info',
        transports: transports,
    })
    loggers.set(label, logger)
    return logger
}

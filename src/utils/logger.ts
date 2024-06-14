import { LogLevels, consola, type ConsolaInstance } from 'consola';

export interface ILogger {
    log(message?: any, ...optionalParams: any[]): void;
    debug(message?: any, ...optionalParams: any[]): void;
    info(message?: any, ...optionalParams: any[]): void;
    warn(message?: any, ...optionalParams: any[]): void;
    error(message?: any, ...optionalParams: any[]): void;
}

export class Logger {
    readonly consola: ConsolaInstance;
    readonly prefix?: string;

    constructor(prefix?: string) {
        this.prefix = prefix?.trim();

        if (this.prefix !== undefined) {
            this.consola = consola.withTag(this.prefix?.trim());
        } else {
            this.consola = consola;
        }

        if (import.meta.dev) {
            this.consola.level = LogLevels.debug;
        } else {
            this.consola.level = LogLevels.info;
        }
    }

    log(message?: any, ...optionalParams: any[]): void {
        this.consola.log(message, ...optionalParams);
    }

    debug(message?: any, ...optionalParams: any[]): void {
        this.consola.debug(message, ...optionalParams);
    }

    info(message?: any, ...optionalParams: any[]): void {
        this.consola.info(message, ...optionalParams);
    }

    warn(message?: any, ...optionalParams: any[]): void {
        this.consola.warn(message, ...optionalParams);
    }

    error(message?: any, ...optionalParams: any[]): void {
        this.consola.error(message, ...optionalParams);
    }
}

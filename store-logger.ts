import { Option } from 'funfix';
import { ILogger, IStoreLogger, IStoreReader, IStoreWriter } from './types';

/**
 * This File:
 *
 * Contains the `StoreLogger` class which is responsible for logging methods
 * called on `IStoreReader` and `IStoreWriter`.
 *
 * This demonstrates the "Single Responsibility Principle" which is the "S" in
 * SOLID.
 *
 * It also maintains the segregation of the interfaces and demonstrates
 * "Implementation Hiding" (a more descriptive term for Encapsulation). This
 * is done by hiding the implementation of the logger behind the interface
 * IStoreLogger which defines the logging methods that can be called.
 *
 * The class also is a decorator for `IStoreReader` and `IStoreWriter` to add
 * debug logging when calling `read` and `save` respectively.
 *
 */
export class StoreLogger implements IStoreReader, IStoreWriter, IStoreLogger {

  public reader: IStoreReader;
  public writer: IStoreWriter;
  public logger: ILogger;
  public context: string;

  constructor(context: string, log: ILogger, reader: IStoreReader, writer: IStoreWriter) {
    this.context = context;
    this.logger = log;
    this.reader = reader;
    this.writer = writer;
  }

  public read(id: number): Option<string> {
    this.log.debug(`  ${this.context}::read() for id ${id}`);
    const value: Option<string> = this.reader.read(id);
    if (value.isEmpty()) {
      this.log.debug(`  ${this.context}::read() found no message.`);
    } else {
      this.log.debug(`  ${this.context}::read() found message: ${value.get()}`);
    }
    return value;
  }

  public save(id: number, message: string): void {
    this.log.debug(`  ${this.context}::save() started for id ${id}`);
    this.writer.save(id, message);
    this.log.debug(`  ${this.context}::save() done for id ${id}`);
  }

  private get log(): ILogger {
    return this.logger;
  }

  public saving(id: number, message: string): void {
    this.log.info(`> ${this.context}::SAVING, ${id}, "${message}"`);
  }

  public savingDone(id: number, message: string): void {
    this.log.info(`< ${this.context}::SAVED, ${id}, "${message}"`);
  }

  public reading(id: number): void {
    this.log.info(`> ${this.context}::READING, ${id}`);
  }

  public readingDone(id: number): void {
    this.log.info(`< ${this.context}::READ, ${id}`);
  }
}

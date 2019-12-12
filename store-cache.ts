import { Option } from 'funfix';
import { IStoreReader, IStoreWriter } from './types';

/**
 * This File:
 *
 * Contains the class `StoreCache` - A cache implementation for a Store.
 *
 * This decorates `IStoreReader` and `IStoreWriter` with an in memory cache.
 */
export class StoreCache implements IStoreReader, IStoreWriter {

  public reader: IStoreReader;
  public writer: IStoreWriter;
  public cache: {[id: number]: string; };

  constructor(writer: IStoreWriter, reader: IStoreReader) {
    this.cache = {};
    this.reader = reader;
    this.writer = writer;
  }

  public read(id: number): Option<string> {
    if (id in this.cache) {
      return Option.of(this.cache[id]);
    }

    const value: Option<string> = this.reader.read(id);

    if (!value.isEmpty()) {
      this.cache[id] = value.get();
    }

    return value;
  }

  public save(id: number, message: string): void {
    this.writer.save(id, message);
    this.cache[id] = message;
  }
}

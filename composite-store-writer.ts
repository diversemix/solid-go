import { IStoreWriter } from './types';

/**
 * This File:
 *
 * Contains the class `CompositeStoreWriter` to be used in one of the exercises.
 *
 */
export class CompositeStoreWriter implements IStoreWriter {

  public writers: IStoreWriter[];

  constructor(writers: IStoreWriter[]) {
    this.writers = writers;
  }

  public save(id: number, message: string): void {
    this.writers.forEach((writer) => writer.save(id, message));
  }

}

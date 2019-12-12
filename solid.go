// Solid example program
package main

import (
	"fmt"

	"github.com/diversemix/solid-go/storage"
)

func main() {
	fs := storage.FileStore{}

	fmt.Println(fs.Read(123))
}

/*
import { FileStore } from './file-store';
import { StoreCache } from './store-cache';
import { StoreLogger } from './store-logger';
import { IStoreLogger, IStoreReader, IStoreWriter } from './types';
*/
/**
 * Welcome to the code for this tutorial.
 *
 * The comments are designed for you to delete as you go, once you've read and
 * understood them. If you want to you can use a package like `strip-comments`
 * to remove them all.
 *
 * Aim:
 *
 * To use an abstract `message-store` to serve as an example of the way we
 * approach coding for the server-side. In general this introduces the reader
 * to the SOLID principles using design patterns.
 */

/**
 * This File:
 *
 * Demonstrates how a message store can be constructed in different way to
 * extend functionality using composition and the decorator pattern. This
 * splits the areas of responsibility out into individual classes thus
 * demonstrating the Single Responsibility Principle (SRP).
 */

/**
 * The following three createMessageStore...() functions create different types
 * of message-store:
 *
 * createMessageStore
 * - Implements using the FileStore.
 *
 * createMessageStoreWithLogging
 * - Implements FileStore and decorated with the StoreLogger.
 *
 * createMessageStoreWithLoggingAndCache
 * - Implements FileStore and decorated with the StoreLogger this is then
 *   further decorated with a cache which itself is decorated with a logger.

interface IStore extends IStoreReader, IStoreWriter {}

const createMessageStore = (): IStore => {
  return new FileStore('');
};

const createMessageStoreWithLogging = (): IStore => {
  const store = new FileStore('');
  return new StoreLogger('FileStore', console, store, store);
};

const createMessageStoreWithLoggingAndCache = (): IStore => {
  const store = new FileStore('');
  const loggingStore = new StoreLogger('FileStore', console, store, store);
  const storeCache = new StoreCache(loggingStore, loggingStore);
  return new StoreLogger('StoreCache', console, storeCache, storeCache);
};
*/

/**
 * runTests() - runs a read-save-read on the store that is passed in. It
 * will also use the logger interface if there is one.

const runTests = (store: IStore, logger: IStoreLogger | null) => {
  const testId = 123;
  const message = 'message1';

  if (logger) { logger.reading(testId); }
  store.read(testId);
  if (logger) { logger.readingDone(testId); }

  if (logger) { logger.saving(testId, message); }
  store.save(testId, message);
  if (logger) { logger.savingDone(testId, message); }

  if (logger) { logger.reading(testId); }
  store.read(testId);
  if (logger) { logger.readingDone(testId); }
};
*/
/**
const run = (store: any) => {
   * This if statement is here because the first example does not implement
   * the IStoreLogger interface.
  if (store instanceof FileStore) {
    runTests(store, null);
  } else {
    runTests(store as IStore, store as IStoreLogger);
  }
};
*/

/**
 * Main section of the application which runs the tests with differently
 * constructed message-stores.
const log = console;

log.log('---- createMessageStore');
run(createMessageStore());

log.log('---- createMessageStoreWithLogging');
run(createMessageStoreWithLogging());

log.log('---- createMessageStoreWithLoggingAndCache');
run(createMessageStoreWithLoggingAndCache());
*/

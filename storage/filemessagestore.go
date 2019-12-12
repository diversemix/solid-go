// Package storage for all your storage needs
package storage

/**
 * This File:
 *
 * Contains the class `FileStore` - An implementation that uses the file system.
 *
 * This demonstrates the "Interface Segregation Principle" which is the "I" in
 * SOLID. Note that we could easily define an `IStore` interface in `types.ts`
 * which would extend both `IStoreReader` and `IStoreWriter` - but this would be
 * premature as this would constrain the way the `FileStore` class could be
 * used too early. This constraint is part of `index.ts` which does indeed
 * define this interface `IStore` for convenience.
 */

// FileStore - stores messages in file system
type FileStore struct {
	cwd string
}

func (r FileStore) Read(id int) (string, error) {
	// TODO: open file and read message
	return "message", nil
}

// Save to file
func (r FileStore) Save(id int, message string) {
	// TODO: write message to file
}

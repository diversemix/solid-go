// Package mstypes for the message store
package mstypes

// IStoreReader - for all things that can read messages
type IStoreReader interface {
	Read(id int) (string, error)
}

// IStoreWriter - for all things that can write messages
type IStoreWriter interface {
	Save(id int, message string)
}

// ILogger - for all things capable of logging
type ILogger interface {
	info(message string)
	error(message string)
}

// IStoreLogger - for all things that can cache
type IStoreLogger interface {
	saving(id int, message string)
	savingDone(id int, message string)
	reading(id int)
	readingDone(id int)
}

// ICacheLogger - for all things that cache logging
type ICacheLogger interface {
	cacheHit(id int, message string)
	savingMiss(id int)
}

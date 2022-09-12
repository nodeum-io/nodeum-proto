package storage

import (
	"context"
	"io"
	"time"
)

type HandlerOptions map[string]interface{}

type ReaddirFunc func(info NodeInfo)

type Handler interface {
	// Prepare is called when the Handler is needed for the first time
	Prepare(ctx context.Context) error
	// Dispose is called when the Handler is not needed anymore
	Dispose(ctx context.Context) error

	// Storage returns the storage sets with `NewStorageHandler`
	Storage() Storage
	// Options returns the storage sets with `NewStorageHandler`
	Options() HandlerOptions

	// Remove removes a node
	Remove(ctx context.Context, path string) error
	// Readdir iterates on every node inside the folder
	Readdir(ctx context.Context, path string, readdirFn ReaddirFunc) error

	// NodeInfo returns information for the node located at `path`
	NodeInfo(ctx context.Context, path string) (NodeInfo, error)
	// SetNodeInfo sets info like FileMode, UID, GID and times. Returns info effectively set.
	SetNodeInfo(ctx context.Context, path string, info NodeInfo) (NodeInfo, error)

	Reader(ctx context.Context, path string) (io.Reader, error)
	Writer(ctx context.Context, path string, opts ...WriterOption) (io.Writer, error)

	// Finalize is called after each operation
	Finalize(ctx context.Context) error
}

type Expirable interface {
	// Expires returns the expiration date.
	// Return zero time if it doesn't expires.
	Expires() time.Time
}

// DirCreator is an optional interface allowing folder creation
type DirCreator interface {
	// Mkdir creates a directory at `path`. Also create parents if missing.
	// Returns infos about all folders created, with the last item being `path`.
	Mkdir(ctx context.Context, path string) ([]NodeInfo, error)
}

type HandlerProvider interface {
	NewStorageHandler(Storage, HandlerOptions) Handler
}

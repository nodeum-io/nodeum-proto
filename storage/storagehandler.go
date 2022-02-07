package storage

import (
	"io"

	"github.com/nodeum-io/nodeum-plugins/protobuf/types/nodeinfopb"
	"github.com/nodeum-io/nodeum-plugins/protobuf/types/storagespb"
)

type HandlerOptions map[string]interface{}

type ReaddirFunc func(info *nodeinfopb.NodeInfo)

type Handler interface {
	// Prepare is called when the Handler is needed for the first time
	Prepare() error
	// Dispose is called when the Handler is not needed anymore
	Dispose() error

	// Storage returns the storage sets with `NewStorageHandler`
	Storage() *storagespb.Storage
	// Options returns the storage sets with `NewStorageHandler`
	Options() HandlerOptions

	// Remove removes a node
	Remove(path string) error
	// Readdir iterates on every node inside the folder
	Readdir(path string, readdirFn ReaddirFunc) error

	// NodeInfo returns information for the node located at `path`
	NodeInfo(path string) (*nodeinfopb.NodeInfo, error)
	// SetNodeInfo sets info like FileMode, UID, GID and times. Returns info effectively set.
	SetNodeInfo(path string, info *nodeinfopb.NodeInfo) (*nodeinfopb.NodeInfo, error)

	Reader(path string) (io.Reader, error)
	Writer(path string, exclusive bool) (io.Writer, error)

	// Finalize is called after each operation
	Finalize() error
}

// DirCreator is an optional interface allowing folder creation
type DirCreator interface {
	// Mkdir creates a directory at `path`. Also create parents if missing.
	// Returns infos about all folders created, with the last item being `path`.
	Mkdir(path string) ([]*nodeinfopb.NodeInfo, error)
}

type HandlerProvider interface {
	NewStorageHandler(*storagespb.Storage, HandlerOptions) Handler
}

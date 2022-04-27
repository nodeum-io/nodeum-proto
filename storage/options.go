package storage

type WriterOptions struct {
	// Exclusive will fail the writing if the destination already exists
	Exclusive bool
	// NodeInfo indicate metadata that also need to be applied.
	NodeInfo *NodeInfo
}

type WriterOption func(o *WriterOptions)

func WithExclusive(v bool) WriterOption {
	return func(o *WriterOptions) {
		o.Exclusive = v
	}
}

func WithNodeInfo(v *NodeInfo) WriterOption {
	return func(o *WriterOptions) {
		o.NodeInfo = v
	}
}

package storage

type WriterOptions struct {
	// Exclusive will fail the writing if the destination already exists
	Exclusive bool
	// Size indicate the total size that is expected to be copied
	Size int64
}

type WriterOption func(o *WriterOptions)

func WithExclusive(v bool) WriterOption {
	return func(o *WriterOptions) {
		o.Exclusive = v
	}
}

func WithSize(v int64) WriterOption {
	return func(o *WriterOptions) {
		o.Size = v
	}
}

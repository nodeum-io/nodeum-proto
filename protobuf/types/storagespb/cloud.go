package storagespb

// IsS3FSCompatible checks if the provider is s3fs compatible
func (p Bucket_Provider) IsS3FSCompatible() bool {
	return p != Bucket_UNDEFINED
}

// IsS3ProxyCompatible checks if the provider needs S3Proxy
func (p Bucket_Provider) IsS3ProxyCompatible() bool {
	return p == Bucket_AZURE || p == Bucket_OPENSTACK_SWIFT
}

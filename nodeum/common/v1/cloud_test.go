package commonv1

import (
	"testing"
)

func TestBucket_Provider_IsS3FSCompatible(t *testing.T) {
	tests := []struct {
		name string
		p    Bucket_Provider
		want bool
	}{
		{"Bucket_PROVIDER_UNSPECIFIED", Bucket_PROVIDER_UNSPECIFIED, false},
		{"Bucket_PROVIDER_GENERIC_S3", Bucket_PROVIDER_GENERIC_S3, true},
		{"Bucket_PROVIDER_AMAZON_AWS_S3", Bucket_PROVIDER_AMAZON_AWS_S3, true},
		{"Bucket_PROVIDER_CLOUDIAN_HYPERSTORE", Bucket_PROVIDER_CLOUDIAN_HYPERSTORE, true},
		{"Bucket_PROVIDER_SCALITY_RING", Bucket_PROVIDER_SCALITY_RING, true},
		{"Bucket_PROVIDER_DELL_EMC_ECS", Bucket_PROVIDER_DELL_EMC_ECS, true},
		{"Bucket_PROVIDER_AZURE", Bucket_PROVIDER_AZURE, true},
		{"Bucket_PROVIDER_GOOGLE_CLOUD_STORAGE", Bucket_PROVIDER_GOOGLE_CLOUD_STORAGE, true},
		{"Bucket_PROVIDER_OPENSTACK_SWIFT", Bucket_PROVIDER_OPENSTACK_SWIFT, true},
		{"Bucket_PROVIDER_WASABI", Bucket_PROVIDER_WASABI, true},
		{"Bucket_PROVIDER_QUANTUM_ACTIVE_SCALE", Bucket_PROVIDER_QUANTUM_ACTIVE_SCALE, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsS3FSCompatible(); got != tt.want {
				t.Errorf("Bucket_Provider.IsS3FSCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_Provider_IsS3ProxyCompatible(t *testing.T) {
	tests := []struct {
		name string
		p    Bucket_Provider
		want bool
	}{
		{"Bucket_PROVIDER_UNSPECIFIED", Bucket_PROVIDER_UNSPECIFIED, false},
		{"Bucket_PROVIDER_GENERIC_S3", Bucket_PROVIDER_GENERIC_S3, false},
		{"Bucket_PROVIDER_AMAZON_AWS_S3", Bucket_PROVIDER_AMAZON_AWS_S3, false},
		{"Bucket_PROVIDER_CLOUDIAN_HYPERSTORE", Bucket_PROVIDER_CLOUDIAN_HYPERSTORE, false},
		{"Bucket_PROVIDER_SCALITY_RING", Bucket_PROVIDER_SCALITY_RING, false},
		{"Bucket_PROVIDER_DELL_EMC_ECS", Bucket_PROVIDER_DELL_EMC_ECS, false},
		{"Bucket_PROVIDER_AZURE", Bucket_PROVIDER_AZURE, true},
		{"Bucket_PROVIDER_GOOGLE_CLOUD_STORAGE", Bucket_PROVIDER_GOOGLE_CLOUD_STORAGE, false},
		{"Bucket_PROVIDER_OPENSTACK_SWIFT", Bucket_PROVIDER_OPENSTACK_SWIFT, true},
		{"Bucket_PROVIDER_WASABI", Bucket_PROVIDER_WASABI, false},
		{"Bucket_PROVIDER_QUANTUM_ACTIVE_SCALE", Bucket_PROVIDER_QUANTUM_ACTIVE_SCALE, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsS3ProxyCompatible(); got != tt.want {
				t.Errorf("Bucket_Provider.IsS3ProxyCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

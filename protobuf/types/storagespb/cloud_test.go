package storagespb

import (
	"testing"
)

func TestBucket_Provider_IsS3FSCompatible(t *testing.T) {
	tests := []struct {
		name string
		p    Bucket_Provider
		want bool
	}{
		{"Bucket_UNDEFINED", Bucket_UNDEFINED, false},
		{"Bucket_GENERIC_S3", Bucket_GENERIC_S3, true},
		{"Bucket_AMAZON_AWS_S3", Bucket_AMAZON_AWS_S3, true},
		{"Bucket_CLOUDIAN_HYPERSTORE", Bucket_CLOUDIAN_HYPERSTORE, true},
		{"Bucket_SCALITY_RING", Bucket_SCALITY_RING, true},
		{"Bucket_DELL_EMC_ECS", Bucket_DELL_EMC_ECS, true},
		{"Bucket_AZURE", Bucket_AZURE, true},
		{"Bucket_GOOGLE_CLOUD_STORAGE", Bucket_GOOGLE_CLOUD_STORAGE, true},
		{"Bucket_OPENSTACK_SWIFT", Bucket_OPENSTACK_SWIFT, true},
		{"Bucket_WASABI", Bucket_WASABI, true},
		{"Bucket_QUANTUM_ACTIVE_SCALE", Bucket_QUANTUM_ACTIVE_SCALE, true},
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
		{"Bucket_UNDEFINED", Bucket_UNDEFINED, false},
		{"Bucket_GENERIC_S3", Bucket_GENERIC_S3, false},
		{"Bucket_AMAZON_AWS_S3", Bucket_AMAZON_AWS_S3, false},
		{"Bucket_CLOUDIAN_HYPERSTORE", Bucket_CLOUDIAN_HYPERSTORE, false},
		{"Bucket_SCALITY_RING", Bucket_SCALITY_RING, false},
		{"Bucket_DELL_EMC_ECS", Bucket_DELL_EMC_ECS, false},
		{"Bucket_AZURE", Bucket_AZURE, true},
		{"Bucket_GOOGLE_CLOUD_STORAGE", Bucket_GOOGLE_CLOUD_STORAGE, false},
		{"Bucket_OPENSTACK_SWIFT", Bucket_OPENSTACK_SWIFT, true},
		{"Bucket_WASABI", Bucket_WASABI, false},
		{"Bucket_QUANTUM_ACTIVE_SCALE", Bucket_QUANTUM_ACTIVE_SCALE, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsS3ProxyCompatible(); got != tt.want {
				t.Errorf("Bucket_Provider.IsS3ProxyCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

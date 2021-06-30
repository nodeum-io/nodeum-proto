package storage

import "testing"

func TestCloudProvider_IsS3FSCompatible(t *testing.T) {
	tests := []struct {
		name string
		p    CloudProvider
		want bool
	}{
		{"GenericS3", GenericS3, true},
		{"AmazonAwsS3", AmazonAwsS3, true},
		{"CloudianHyperstore", CloudianHyperstore, true},
		{"ScalityRing", ScalityRing, true},
		{"DellEmcEcs", DellEmcEcs, true},
		{"Azure", Azure, true},
		{"GoogleCloudStorage", GoogleCloudStorage, true},
		{"OpenstackSwift", OpenstackSwift, true},
		{"Wasabi", Wasabi, true},
		{"QuantumActivescale", QuantumActivescale, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsS3FSCompatible(); got != tt.want {
				t.Errorf("CloudProvider.IsS3FSCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudProvider_IsS3ProxyCompatible(t *testing.T) {
	tests := []struct {
		name string
		p    CloudProvider
		want bool
	}{
		{"GenericS3", GenericS3, false},
		{"AmazonAwsS3", AmazonAwsS3, false},
		{"CloudianHyperstore", CloudianHyperstore, false},
		{"ScalityRing", ScalityRing, false},
		{"DellEmcEcs", DellEmcEcs, false},
		{"Azure", Azure, true},
		{"GoogleCloudStorage", GoogleCloudStorage, false},
		{"OpenstackSwift", OpenstackSwift, true},
		{"Wasabi", Wasabi, false},
		{"QuantumActivescale", QuantumActivescale, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsS3ProxyCompatible(); got != tt.want {
				t.Errorf("CloudProvider.IsS3ProxyCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

package commonv1

import (
	"testing"
)

func TestNASShare_Protocol_IsNFS(t *testing.T) {
	tests := []struct {
		name string
		n    NASShare_Protocol
		want bool
	}{
		{"NASShare_PROTOCOL_UNSPECIFIED", NASShare_PROTOCOL_UNSPECIFIED, false},
		{"NASShare_PROTOCOL_SMB_V1", NASShare_PROTOCOL_SMB_V1, false},
		{"NASShare_PROTOCOL_NFS_V3", NASShare_PROTOCOL_NFS_V3, true},
		{"NASShare_PROTOCOL_STORE_NEXT_V5", NASShare_PROTOCOL_STORE_NEXT_V5, false},
		{"NASShare_PROTOCOL_NFS_V4", NASShare_PROTOCOL_NFS_V4, true},
		{"NASShare_PROTOCOL_SMB_V2_1", NASShare_PROTOCOL_SMB_V2_1, false},
		{"NASShare_PROTOCOL_SMB_V3", NASShare_PROTOCOL_SMB_V3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsNFS(); got != tt.want {
				t.Errorf("NASShare_Protocol.IsNFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNASShare_Protocol_IsCIFS(t *testing.T) {
	tests := []struct {
		name string
		n    NASShare_Protocol
		want bool
	}{
		{"NASShare_PROTOCOL_UNSPECIFIED", NASShare_PROTOCOL_UNSPECIFIED, false},
		{"NASShare_PROTOCOL_SMB_V1", NASShare_PROTOCOL_SMB_V1, true},
		{"NASShare_PROTOCOL_NFS_V3", NASShare_PROTOCOL_NFS_V3, false},
		{"NASShare_PROTOCOL_STORE_NEXT_V5", NASShare_PROTOCOL_STORE_NEXT_V5, false},
		{"NASShare_PROTOCOL_NFS_V4", NASShare_PROTOCOL_NFS_V4, false},
		{"NASShare_PROTOCOL_SMB_V2_1", NASShare_PROTOCOL_SMB_V2_1, true},
		{"NASShare_PROTOCOL_SMB_V3", NASShare_PROTOCOL_SMB_V3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsCIFS(); got != tt.want {
				t.Errorf("NASShare_Protocol.IsCIFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

package storage

import "testing"

func TestNASProtocol_IsNFS(t *testing.T) {
	tests := []struct {
		name string
		n    NASProtocol
		want bool
	}{
		{"UndefinedNASProtocol", UndefinedNASProtocol, false},
		{"SMBv1", SMBv1, false},
		{"NFSv3", NFSv3, true},
		{"StoreNextV5", StoreNextV5, false},
		{"NFSv4", NFSv4, true},
		{"SMBv21", SMBv21, false},
		{"SMBv3", SMBv3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsNFS(); got != tt.want {
				t.Errorf("NASProtocol.IsNFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNASProtocol_IsCIFS(t *testing.T) {
	tests := []struct {
		name string
		n    NASProtocol
		want bool
	}{
		{"UndefinedNASProtocol", UndefinedNASProtocol, false},
		{"SMBv1", SMBv1, true},
		{"NFSv3", NFSv3, false},
		{"StoreNextV5", StoreNextV5, false},
		{"NFSv4", NFSv4, false},
		{"SMBv21", SMBv21, true},
		{"SMBv3", SMBv3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsCIFS(); got != tt.want {
				t.Errorf("NASProtocol.IsCIFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

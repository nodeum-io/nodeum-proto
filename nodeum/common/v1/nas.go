package commonv1

// IsNFS checks if the type is NFS compatible
func (n NASShare_Protocol) IsNFS() bool {
	return n == NASShare_PROTOCOL_NFS_V3 || n == NASShare_PROTOCOL_NFS_V4
}

// IsCIFS checks if the type is CIFS (samba) compatible
func (n NASShare_Protocol) IsCIFS() bool {
	return n == NASShare_PROTOCOL_SMB_V1 ||
		n == NASShare_PROTOCOL_SMB_V2_1 ||
		n == NASShare_PROTOCOL_SMB_V3
}

// Version returns a string of the numerical version
func (n NASShare_Protocol) Version() string {
	switch n {
	case NASShare_PROTOCOL_NFS_V4:
		return "4.2"
	case NASShare_PROTOCOL_NFS_V3:
		return "3"
	case NASShare_PROTOCOL_SMB_V3:
		return "3.0"
	case NASShare_PROTOCOL_SMB_V2_1:
		return "2.1"
	case NASShare_PROTOCOL_SMB_V1:
		return "1.0"
	}
	return ""
}

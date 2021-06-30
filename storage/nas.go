package storage

// NASProtocol enumeration for NAS Share
type NASProtocol uint16

const (
	SMBv1 NASProtocol = iota
	NFSv3
	StoreNextV5
	NFSv4
	SMBv21
	SMBv3
)

// IsNFS checks if the type is NFS compatible
func (n NASProtocol) IsNFS() bool {
	return n == NFSv3 || n == NFSv4
}

// IsCIFS checks if the type is CIFS (samba) compatible
func (n NASProtocol) IsCIFS() bool {
	return n == SMBv1 ||
		n == SMBv21 ||
		n == SMBv3
}

// Version returns a string of the numerical version
func (n NASProtocol) Version() string {
	switch n {
	case NFSv4:
		return "4.2"
	case NFSv3:
		return "3"
	case SMBv3:
		return "3.0"
	case SMBv21:
		return "2.1"
	case SMBv1:
		return "1.0"
	}
	return ""
}

type NASShare struct {
	MID          uint64
	MName        string
	MNASName     string
	MPrimaryName string

	Protocol NASProtocol
	Host     string
	Path     string
	Username string
	Password string
	Options  string
}

func (b NASShare) ID() interface{} {
	return b.MID
}
func (b NASShare) Type() Type {
	return NASShareType
}
func (b NASShare) Name() string {
	return b.MName
}
func (b NASShare) ParentName() string {
	return b.MNASName
}
func (b NASShare) PrimaryName() string {
	return b.MPrimaryName
}

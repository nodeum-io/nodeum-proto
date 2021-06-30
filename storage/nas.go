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

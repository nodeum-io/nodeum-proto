package storage

// Type is an enumeration of type of storages
type Type string

// Type is an enumeration of type of storages
const (
	ContainerType   Type = "container"
	NASShareType    Type = "nas-share"
	CloudBucketType Type = "cloud-bucket"
)

type Storage interface {
	ID() interface{}
	Type() Type
	Name() string
	ParentName() string
	PrimaryName() string
}

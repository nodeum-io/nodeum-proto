package nodeumplugins

// StorageType is an enumeration of type of storages
type StorageType string

// Type is an enumeration of type of storages
const (
	ContainerType   StorageType = "container"
	NASShareType    StorageType = "nas-share"
	CloudBucketType StorageType = "cloud-bucket"
)

type Storage interface {
	ID() interface{}
	Type() StorageType
	Name() string
	ParentName() string
	PrimaryName() string
}

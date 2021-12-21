package storage

import "fmt"

// CloudProvider enumeration for Cloud connector
type CloudProvider uint16

const (
	UndefinedCloudProvider CloudProvider = iota
	GenericS3
	AmazonAwsS3
	CloudianHyperstore
	ScalityRing
	DellEmcEcs
	Azure
	GoogleCloudStorage
	OpenstackSwift
	Wasabi
	QuantumActivescale
)

// IsS3FSCompatible checks if the provider is s3fs compatible
func (p CloudProvider) IsS3FSCompatible() bool {
	return p != UndefinedCloudProvider
}

// IsS3ProxyCompatible checks if the provider needs S3Proxy
func (p CloudProvider) IsS3ProxyCompatible() bool {
	return p == Azure || p == OpenstackSwift
}

type Bucket struct {
	MID            uint64
	MName          string
	MConnectorName string
	MPrimaryName   string

	URL       string
	Provider  *CloudProvider
	Region    string
	AccessKey string
	SecretKey string
	Options   string
}

func (b Bucket) ID() interface{} {
	return b.MID
}
func (b Bucket) Type() Type {
	return CloudBucketType
}
func (b Bucket) Name() string {
	return b.MName
}
func (b Bucket) ParentName() string {
	return b.MConnectorName
}
func (b Bucket) PrimaryName() string {
	return b.MPrimaryName
}

func (b Bucket) String() string {
	return fmt.Sprintf("%v[%v;#%v]", CloudBucketType, b.MName, b.MID)
}

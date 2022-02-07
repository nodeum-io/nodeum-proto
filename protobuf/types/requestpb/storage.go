package requestpb

import storagespb "github.com/nodeum-io/nodeum-plugins/protobuf/types/storagespb"

func (req *Request) GetSourceStorage() *storagespb.Storage {
	switch reqt := req.GetRequest().(type) {
	case *Request_Copy:
		return reqt.Copy.GetSource()
	case *Request_Remove:
		return reqt.Remove.GetStorage()
	case *Request_Scan:
		return reqt.Scan.GetStorage()
	}

	return nil
}

func (req *Request) GetDestinationStorage() *storagespb.Storage {
	switch reqt := req.GetRequest().(type) {
	case *Request_Copy:
		return reqt.Copy.GetDestination()
	}

	return nil
}

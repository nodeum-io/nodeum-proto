package commonv1

func (req *Request) GetSourceStorage() *Storage {
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

func (req *Request) GetDestinationStorage() *Storage {
	switch reqt := req.GetRequest().(type) {
	case *Request_Copy:
		return reqt.Copy.GetDestination()
	}

	return nil
}

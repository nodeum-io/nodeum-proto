package commonv1

func (req *Request) GetSourcePath() string {
	switch reqt := req.GetRequest().(type) {
	case *Request_Copy:
		return reqt.Copy.GetSourcePath()
	case *Request_Remove:
		return reqt.Remove.GetPath()
	case *Request_Scan:
		return reqt.Scan.GetPath()
	}

	return ""
}

func (req *Request) GetDestinationPath() string {
	switch reqt := req.GetRequest().(type) {
	case *Request_Copy:
		return reqt.Copy.GetDestinationPath()
	}

	return ""
}

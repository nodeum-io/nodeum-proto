package commonv1

func (s *Storage) ID() (id uint64) {
	switch st := s.GetStorage().(type) {
	case *Storage_Bucket:
		id = st.Bucket.GetId()
	case *Storage_NasShare:
		id = st.NasShare.GetId()
	case *Storage_Container:
		id = st.Container.GetId()
	}
	return id
}

func (s *Storage) Name() (name string) {
	switch st := s.GetStorage().(type) {
	case *Storage_Bucket:
		name = st.Bucket.GetName()
	case *Storage_NasShare:
		name = st.NasShare.GetName()
	case *Storage_Container:
		name = st.Container.GetName()
	}
	return name
}

func (s *Storage) ParentName() (name string) {
	switch st := s.GetStorage().(type) {
	case *Storage_Bucket:
		name = st.Bucket.GetConnectorName()
	case *Storage_NasShare:
		name = st.NasShare.GetNasName()
	}
	return name
}

func (s *Storage) PrimaryName() (name string) {
	switch st := s.GetStorage().(type) {
	case *Storage_Bucket:
		name = st.Bucket.GetPrimaryName()
	case *Storage_NasShare:
		name = st.NasShare.GetPrimaryName()
	case *Storage_Container:
		name = st.Container.GetName()
	}
	return name
}

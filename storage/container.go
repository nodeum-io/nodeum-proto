package storage

type Container struct {
	MID   uint64
	MName string
}

func (b Container) ID() interface{} {
	return b.MID
}
func (b Container) Type() Type {
	return ContainerType
}
func (b Container) Name() string {
	return b.MName
}
func (b Container) ParentName() string {
	return ""
}
func (b Container) PrimaryName() string {
	return b.MName
}

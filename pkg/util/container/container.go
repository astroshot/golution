package container

// Container interface
type Container interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []interface{}
}

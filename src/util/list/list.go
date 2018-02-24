package list

// List interface
type List interface {
	Get(index int) (interface{}, bool)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Swap(pos1, pos2 int)
	Insert(index int, values ...interface{})
}

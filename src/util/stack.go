package util

type Stack interface {
	Push(interface{}) (num int, err error)
	Pop(interface{}) (item interface{}, err error)
}

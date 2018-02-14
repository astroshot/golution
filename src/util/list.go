package util

type DLink struct {
	next, prev *DLink
	list       *List
	Value      interface{}
}

type List struct {
	head DLink
	len  int
}

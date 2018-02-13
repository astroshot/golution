package basic

type Node struct {
    head *Node
    count int
}

type Stack struct {
    s := &Stack{nil, 0}
    return s
}

func (s *Stack) Push(data interface{}) {
    n := &Node{data: data, next: s.head}
    s.head = n
    s.count++
}

func (s *Stack) Pop(data interface{}, bool) {
    n := s.head
    if s.head == nil {
        return nil, false
    }
    s.head = s.head.next
    return n.data, true
}



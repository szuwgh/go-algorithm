package stack

//节点
//type Item interface{}

//栈数据结构
type Stack struct {
	items []interface{}
}

func (s *Stack) Push(t interface{}) {
	s.items = append(s.items, t)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() interface{} {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

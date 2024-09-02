package queue

//栈数据结构
type Queue struct {
	items []interface{}
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Put(t interface{}) {
	q.items = append(q.items, t)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Get() interface{} {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item
}

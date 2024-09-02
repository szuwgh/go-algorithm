package linklist

import "testing"

func Test_LinkList(t *testing.T) {
	node := &LinkNode{0, nil}
	node.Add(1)
	node.Add(2)
	node.Add(3)
	node.Add(4)
	node.Add(5)
	node.Add(6)
	node.Add(7)
	node.Print()
}

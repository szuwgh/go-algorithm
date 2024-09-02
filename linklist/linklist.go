package linklist

import "fmt"

type Item interface{}

type LinkNode struct {
	Value int
	Next  *LinkNode
}

//添加
func (node *LinkNode) Add(Value int) {
	point := node
	for point.Next != nil {
		point = point.Next
	}
	newNode := &LinkNode{Value, nil}
	point.Next = newNode
}

func (node *LinkNode) Length() int {

	iterator := node
	var length int
	for iterator.Next != nil {
		length++
		iterator = iterator.Next
	}
	return length
}

func (node *LinkNode) Insert(index int, Value int) {
	length := node.Length()
	if index < 0 || index > length {
		return
	}
	point := node
	for i := 0; i < index; i++ {
		point = node.Next
	}
	newNode := &LinkNode{Value, nil}
	newNode.Next = point.Next
	point.Next = newNode
}

func (node *LinkNode) Delete(index int) {
	length := node.Length()
	if index < 0 || index > length {
		return
	}
}

func (node *LinkNode) Print() {
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
}

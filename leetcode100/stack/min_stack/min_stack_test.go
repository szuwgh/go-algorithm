package minstack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin()) // 返回 -3
	minStack.Pop()
	println(minStack.Top())    // 返回 0
	println(minStack.GetMin()) // 返回 -2
}

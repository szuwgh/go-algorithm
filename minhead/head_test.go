package minheap

import (
	"fmt"
	"testing"
)

//4 7 2 5 6 1 0 3 8
func Test_head(t *testing.T) {
	head := NewHeap(9)
	head.Add(4)
	head.Add(7)
	head.Add(2)
	head.Add(5)
	head.Add(6)
	head.Add(1)
	head.Add(0)
	head.Add(3)
	head.Add(8)
	head.Print()
	fmt.Println(head.Get())
}

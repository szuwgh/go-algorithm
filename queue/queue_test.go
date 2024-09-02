package queue

import (
	"fmt"
	"testing"
)

// 结论：s = s[low : high : max] 切片的三个参数的切片截取的意义为 low为截取的起始下标（含）， high为窃取的结束下标（不含high）
func Test_queue(t *testing.T) {
	var queue Queue
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	queue.Put(4)
	queue.Put(5)
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
}

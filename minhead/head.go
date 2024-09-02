package minheap

import (
	"fmt"
)

// 最小堆数据结构
// data 存储堆元素
// Cmp 元素比较函数
type Heap struct {
	data []int
}

type HeapCmpFunc func(interface{}, interface{}) int

func NewHeap(cap int) Heap {
	return Heap{
		data: make([]int, 0, cap),
	}
}

func (heap *Heap) Len() int {
	return len(heap.data)
}

func (heap *Heap) Print() {
	for i, v := range heap.data {
		fmt.Println(i, v)
	}
}

// 向堆中添加新元素
func (heap *Heap) Add(data int) bool {
	idx := heap.Len()
	if idx >= cap(heap.data) {
		heap.scale()
	}
	heap.data = append(heap.data, data)
	heap.shiftUp(idx)
	return true
}

// 获取堆顶元素值与来源标记
func (heap *Heap) Get() (data int) {
	if heap.Len() < 1 {
		return 0
	}

	data = heap.data[0]
	heap.data = heap.data[1:]
	heap.heapify()
	return
}

func (heap *Heap) heapify() {
	firstParent := (heap.Len() - 1) / 2
	for i := firstParent; i >= 0; i-- {
		heap.shiftDown(i)
	}
}

func (heap *Heap) shiftUp(idx int) {
	for idx > 0 {
		if heap.data[idx] < heap.data[parent(idx)] {
			heap.swap(idx, parent(idx))
			idx = parent(idx)
		} else {
			break
		}
	}
}

func (heap *Heap) shiftDown(idx int) {
	l, r := left(idx), right(idx)
	if r < heap.Len() && heap.data[idx] > heap.data[r] {
		heap.swap(idx, r)
	}
	if l < heap.Len() && heap.data[idx] > heap.data[l] {
		heap.swap(idx, l)
	}
}
func (heap *Heap) swap(i, j int) {
	heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
}

func (heap *Heap) scale() {
	cap := len(heap.data) * 2
	if cap == 0 {
		cap = 8
	}
	data := make([]int, len(heap.data), cap)
	copy(data, heap.data)
	heap.data = data
}

func parent(idx int) int {
	return (idx - 1) / 2
}
func left(idx int) int {
	return 2*idx + 1
}
func right(idx int) int {
	return 2*idx + 2
}

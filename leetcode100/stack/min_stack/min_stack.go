package minstack

import "math"

// 最小栈

// 提示
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// 实现 MinStack 类:

// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

// 示例 1:

// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// 输出：
// [null,null,null,null,-3,null,0,-2]

// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

// 关键点解析：
// stack 栈：存储所有的值。
// minStack 栈：存储当前的最小值，确保在每次 push 操作后，minStack 的栈顶都是当前 stack 中的最小值。
// Push 操作：每次压入新值时，同时更新 minStack，将当前的最小值压入 minStack。
// Pop 操作：从两个栈中都弹出栈顶元素。
// Top 操作：只返回 stack 栈的栈顶元素。
// GetMin 操作：直接返回 minStack 的栈顶元素，这样可以在 O(1) 时间内获得最小值。

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	// 初始化时，往minStack里压入一个最大值，便于比较
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	// 将val压入stack
	this.stack = append(this.stack, val)
	// 将当前最小值和val比较，压入较小的那个到minStack
	min := this.minStack[len(this.minStack)-1]
	if val < min {
		min = val
	}
	this.minStack = append(this.minStack, min)
}

func (this *MinStack) Pop() {
	// 弹出stack和minStack栈顶的值
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	// 返回stack的栈顶值
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	// 返回minStack的栈顶值（即当前的最小值）
	return this.minStack[len(this.minStack)-1]
}

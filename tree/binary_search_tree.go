package tree

import (
	"fmt"
)

//节点 二叉树搜索树
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Value = v
}

//插入
func (node *TreeNode) Insert(value int) {
	n := &TreeNode{value, nil, nil}
	node.insertNode(n)
}

func (node *TreeNode) insertNode(newNode *TreeNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.insertNode(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.insertNode(newNode)
		}
	}
}

//搜索
func (node *TreeNode) Search(value int) bool {
	return search(node, value)
}

func search(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return search(node.Left, value)
	}
	if value > node.Value {
		return search(node.Right, value)
	}
	return true
}

//前序遍历
func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}

//中序遍历
func (node *TreeNode) MiddleOrder() {
	if node == nil {
		return
	}
	node.Left.MiddleOrder()
	node.Print()
	node.Right.MiddleOrder()
}

//后序遍历
func (node *TreeNode) PostOrder() {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	node.Print()
}

//蛇形遍历
func (node *TreeNode) zigzagLevelOrder() {
	var data [][]int
	zigzagLevelOrder(node, &data, 0)
	for _, v := range data {
		fmt.Println(v)
	}
}

func zigzagLevelOrder(node *TreeNode, data *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*data) <= level {
		*data = append(*data, []int{})
	}
	curData := (*data)[level]
	if level&1 == 0 {
		curData = append(curData, node.Value)
	} else {
		if len(curData) >= 1 {
			curData = append(curData[:1], curData...)
			curData[0] = node.Value
		} else {
			curData = append(curData, node.Value)
		}
	}
	(*data)[level] = curData
	zigzagLevelOrder(node.Left, data, level+1)
	zigzagLevelOrder(node.Right, data, level+1)
}

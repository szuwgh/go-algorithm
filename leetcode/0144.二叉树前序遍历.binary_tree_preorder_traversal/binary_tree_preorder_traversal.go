package binary_tree_preorder_traversal

import (
	"stack"
	"tree"
)

//前序遍历
type treeNode = tree.TreeNode

func preorderTraversal(root *treeNode) []int {
	var result []int
	stack := &stack.Stack{}
	var top *treeNode
	stack.Push(root)
	for !stack.IsEmpty() {
		top = stack.Pop().(*treeNode)
		result = append(result, top.Value)
		if top.Right != nil {
			stack.Push(top.Right)
		}
		if top.Left != nil {
			stack.Push(top.Left)
		}
	}
	return result
}

package binary_tree_inorder_traversal

import (
	"stack"
	"tree"
)

//中序遍历

type treeNode = tree.TreeNode

func inorderTraversal(root *treeNode) []int {
	var result []int
	stack := &stack.Stack{}
	var top *treeNode
	top = root
	for !stack.IsEmpty() || top != nil {
		for top != nil {
			stack.Push(top)
			top = top.Left
		}
		top = stack.Pop().(*treeNode)
		result = append(result, top.Value)
		top = top.Right
	}
	return result
}

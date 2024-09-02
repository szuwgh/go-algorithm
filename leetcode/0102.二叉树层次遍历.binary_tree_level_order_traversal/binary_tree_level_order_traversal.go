package binary_tree_level_order_traversal

import (
	"queue"
	"tree"
)

//层次遍历
// 题目描述：
// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
type treeNode = tree.TreeNode

func levelOrderTraversal(node *treeNode) [][]int {
	var res [][]int
	queue := &queue.Queue{}
	queue.Put(node)
	var top *treeNode
	for !queue.IsEmpty() {
		size := queue.Size()
		var list []int
		for size > 0 {
			top = queue.Get().(*treeNode)
			list = append(list, top.Value)
			if top.Left != nil {
				queue.Put(top.Left)
			}
			if top.Right != nil {
				queue.Put(top.Right)
			}
			size--
		}
		res = append(res, list)
	}
	return res
}

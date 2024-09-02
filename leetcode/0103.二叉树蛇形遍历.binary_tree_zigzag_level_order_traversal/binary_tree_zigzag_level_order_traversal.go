package binary_tree_zigzag_level_order_traversal

import (
	"queue"
	"tree"
)

//层次遍历
// 题目描述：
// 给定一个二叉树，返回其按蛇形遍历的节点值。

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
//   [20,9],
//   [15,7]
// ]
type treeNode = tree.TreeNode

func zigzagLevelOrderTraversal(node *treeNode) [][]int {
	var res [][]int
	queue := &queue.Queue{}
	queue.Put(node)
	var top *treeNode
	var level int = 0
	for !queue.IsEmpty() {
		size := queue.Size()
		var list []int
		for size > 0 {
			top = queue.Get().(*treeNode)
			list = append(list, top.Value)
			if level&1 == 1 {
				if top.Left != nil {
					queue.Put(top.Left)
				}
				if top.Right != nil {
					queue.Put(top.Right)
				}
			} else {
				if top.Right != nil {
					queue.Put(top.Right)
				}
				if top.Left != nil {
					queue.Put(top.Left)
				}
			}
			size--
		}
		level++
		res = append(res, list)
	}
	return res
}

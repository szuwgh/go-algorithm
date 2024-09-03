package invert_tree

import "fmt"

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

// 示例 1：
//     4
//   /  \
//  2    7
// / \  / \
// 1 3  6  9  =>

//     4
//   /  \
//  7    2
// / \  / \
// 9 6  3  1

// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
// 示例 2：

// 输入：root = [2,1,3]
// 输出：[2,3,1]
// 示例 3：

//  输入：root = []
//  输出：[]

// 代码解释：
// TreeNode 定义了一个二叉树节点的结构。
// invertTree 函数是递归地翻转二叉树的主要逻辑：
// 如果当前节点为空，则直接返回 nil。
// 递归翻转当前节点的左右子树。
// 交换当前节点的左右子树。
// 返回当前节点（即翻转后的树）。
// printTree 函数用于层序遍历打印二叉树的节点值，用于验证结果。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 递归翻转左右子树
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// 交换左右子树
	root.Left = right
	root.Right = left

	return root
}

// 辅助函数，用于打印二叉树（按层序遍历方式）
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			fmt.Printf("%d ", node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			fmt.Printf("nil ")
		}
	}
	fmt.Println()
}

package sorted_array_to_bst

import "fmt"

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
// 平衡
//  二叉搜索树。

// 示例 1：

// //      0
// //    /  \
// //   -3   9
// //  /  \  / \
// // -10    5

// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

// 示例 2：
// //      0
// //    /  \
// //   -10   5
// //   / \   / \
// //    -3      9

// 输入：nums = [1,3]
// 输出：[3,1]
// 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 将一个升序数组转换为平衡二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return buildTree(nums, 0, len(nums)-1)
}

// buildTree 递归构建树
func buildTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree(nums, left, mid-1)
	root.Right = buildTree(nums, mid+1, right)
	return root
}

// printTree 以层序遍历的方式打印树
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				fmt.Print("nil ")
				continue
			}
			fmt.Printf("%d ", node.Val)
			queue = append(queue, node.Left, node.Right)
		}
		fmt.Println()
	}
}

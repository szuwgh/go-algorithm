package diameter_of_binary_tree

// golang 543. 二叉树的直径
// 简单
// 相关标签
// 相关企业
// 给你一棵二叉树的根节点，返回该树的 直径 。

// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

// 两节点之间路径的 长度 由它们之间边数表示。

// 示例 1：

//      1
//    /  \
//   2    3
//  / \  / \
// 4   5

// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 示例 2：

// 输入：root = [1,2]
// 输出：1

// 要解决二叉树的直径问题，我们可以使用递归的方法来计算二叉树的深度，并在计算深度的同时更新二叉树的直径。

// 思路：
// 定义递归函数 depth(node)：这个函数返回以当前节点为根的子树的深度。
// 在计算深度的过程中更新最大直径：在递归计算每个节点的左子树深度 left_depth 和右子树深度 right_depth 时，节点的直径可以定义为 left_depth + right_depth。用一个全局变量 max_diameter 记录最大直径。
// 最终返回直径：在遍历完所有节点后，max_diameter 就是二叉树的直径

// 解释：
// 深度计算：在 depth 函数中，递归计算左右子树的深度，并在回溯时更新最大直径。
// 时间复杂度：该算法遍历了每个节点一次，所以时间复杂度为 O(n)，其中 n 是节点的数量。

// 定义二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// 递归函数计算深度，并更新最大直径
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		// 更新最大直径
		if leftDepth+rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}
		// 返回当前节点的深度
		return 1 + max(leftDepth, rightDepth)
	}

	depth(root)
	return maxDiameter
}

// 辅助函数：返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

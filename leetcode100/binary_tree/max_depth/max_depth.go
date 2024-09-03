package max_depth

// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// 示例 1：

//   3
//  / \
// 9  20
//   /  \
// 15    7

// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
// 示例 2：

// 输入：root = [1,null,2]
// 输出：2

//TreeNode 结构体：定义了一个二叉树节点，包含三个字段：值、左子树和右子树。
//maxDepth 函数：递归计算二叉树的最大深度。对于空节点，深度为 0。对于非空节点，递归计算其左子树和右子树的深度，取最大值并加上 1 表示当前节点的深度。
//max 函数：辅助函数，用于返回两个整数中的较大值。
//main 函数：构建了一个示例二叉树并计算其最大深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 返回左右子树深度的最大值 + 1 (当前节点)
	return max(leftDepth, rightDepth) + 1
}

// 辅助函数，获取两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

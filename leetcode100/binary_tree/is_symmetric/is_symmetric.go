package is_symmetric

// 101. 对称二叉树
// 简单

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

// 示例 1：
//     1
//   /  \
//  2    2
// / \  / \
//3  4  4  3

// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 示例 2：

// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
// 要检查一个二叉树是否轴对称，实际上是判断该二叉树是否是对称的。对称二叉树的定义是：一棵二叉树与它自身的镜像相同。

// 具体思路如下：

// 如果一棵树的根节点为空，则它是对称的。
// 如果根节点不为空，递归地检查根节点的左子树和右子树是否是镜像对称的。
// 左子树的左子节点和右子树的右子节点相等。
// 左子树的右子节点和右子树的左子节点相等。

// 代码解析：
// isSymmetric 函数用于判断树是否对称。它调用 isMirror 函数检查根节点的左右子树是否镜像对称。
// isMirror 函数递归地检查两个子树是否镜像对称：
// 如果两个节点都为空，则它们是对称的。
// 如果只有一个节点为空，则它们不对称。
// 如果两个节点的值相等，递归检查它们的子节点是否镜像对称。
// 复杂度分析：
// 时间复杂度：O(n)，其中 n 是树中的节点数量。每个节点最多被访问一次。
// 空间复杂度：O(h)，其中 h 是树的高度，空间复杂度由递归调用栈的深度决定。

// TreeNode 定义二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 检查二叉树是否轴对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror 检查两个树是否是镜像对称
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

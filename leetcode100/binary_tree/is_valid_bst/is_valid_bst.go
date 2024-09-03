package is_valid_bst

// 验证二叉搜索树
// 中等
// 相关标签
// 相关企业
// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

// 有效 二叉搜索树定义如下：

// 节点的左
// 子树
// 只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// 示例 1：

//    2
//   /  \
//  1    3

// 输入：root = [2,1,3]
// 输出：true
// 示例 2：

//     5
//   /  \
//  1    4
// / \  / \
//      3  6

// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST 函数检查给定的二叉树是否是有效的二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

// isValidBSTHelper 辅助函数，递归检查二叉树
func isValidBSTHelper(node *TreeNode, min *int, max *int) bool {
	// 如果节点为空，则视为有效的BST
	if node == nil {
		return true
	}

	// 如果最大值不为空，且当前节点值大于等于最大值，返回 false
	if max != nil && node.Val >= *max {
		return false
	}

	// 如果最小值不为空，且当前节点值小于等于最小值，返回 false
	if min != nil && node.Val <= *min {
		return false
	}

	// 递归检查右子树，更新最小值为当前节点值
	if !isValidBSTHelper(node.Right, &node.Val, max) {
		return false
	}

	// 递归检查左子树，更新最大值为当前节点值
	if !isValidBSTHelper(node.Left, min, &node.Val) {
		return false
	}

	return true
}

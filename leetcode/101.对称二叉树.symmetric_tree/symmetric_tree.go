package symmetric_tree

import "tree"

// 题目描述：

// 给定一个二叉树，检查它是否是镜像对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

//     1
//    / \
//   2   2
//    \   \
//    3    3
type treeNode = tree.TreeNode

func isSymmetricTree(root *treeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree2(root.Left, root.Right)
}

func isSymmetricTree2(left, right *treeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value && isSymmetricTree2(left.Left, right.Right) && isSymmetricTree2(left.Right, right.Left)
}

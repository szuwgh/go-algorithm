package buildtree

import "fmt"

// 给定两个整数数组 preorder 和 inorder ，
// 其中 preorder 是二叉树的先序遍历，
//  inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

// 示例 1:
//      3
//     /  \
//    9   20
//       /  \
//      15   7

// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// 示例 2:

// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]

// 要从前序遍历和中序遍历构造一棵二叉树，我们可以利用递归的方式来构建。
// 前序遍历的第一个元素是树的根节点，通过在中序遍历中找到这个根节点，
// 我们可以将中序遍历分成左子树和右子树，然后继续递归地构造左右子树。

// 代码讲解：
// 函数 buildTree: 这个函数接收前序遍历和中序遍历的数组，并返回构造好的二叉树的根节点。

// 通过 preorder[0] 获取根节点的值，并找到这个值在 inorder 数组中的索引 rootIndex。
// 将 inorder 数组分割成左子树和右子树的部分。
// 根据左子树 inorder 的长度，将 preorder 数组分割成左子树和右子树的部分。
// 递归地构造左右子树。
// 测试 main 函数: 测试了 preorder 和 inorder 数组，构造了二叉树，并通过前序遍历打印了结果。

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 前序遍历的第一个元素是根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 划分左子树和右子树的中序遍历和前序遍历
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	// 递归构造左右子树
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

// 打印树（前序遍历）用于测试
func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	printTree(node.Left)
	printTree(node.Right)
}

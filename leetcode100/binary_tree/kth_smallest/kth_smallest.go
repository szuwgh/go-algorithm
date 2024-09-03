package kthsmallest

// 二叉搜索树中第 K 小的元素
// 中等
// 提示
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。

// 示例 1：

//     3
//   /  \
//  1    4
// / \
//    2

// 输入：root = [3,1,4,null,2], k = 1
// 输出：1

// 示例 2：

//      5
//     /  \
//    3   6
//   / \
//  2   4
// /
// 1

// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3

// 要在二叉搜索树 (BST) 中找到第 K 小的元素，
// 可以利用中序遍历的特点。中序遍历（左-根-右）
// 可以把 BST 的节点按升序排列，
// 因此我们只需执行中序遍历，
// 直到到达第 K 个节点为止。

// 代码解释：
// TreeNode 定义了二叉树的节点结构，每个节点包含一个整数值 Val 和两个指向其左右子节点的指针 Left 和 Right。

// kthSmallest 函数实现查找第 K 小的元素。我们在函数中定义了一个匿名函数 inorderTraversal，
// 它通过递归实现了对 BST 的中序遍历，并将遍历到的节点值保存到 result 切片中。

// 中序遍历完成后，result[k-1] 即为第 K 小的元素。

// 优化方案：
// 由于题目只要求找到第 K 小的元素，其实在遍历过程中我们可以不必将所有节点存储到 result 中，
// 而是在遍历过程中直接计数，并在计数到第 K 个节点时返回对应的值。下面是优化后的代码：

// TreeNode 定义二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest 查找 BST 中第 K 小的元素
func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int
	var inorderTraversal func(node *TreeNode)

	inorderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderTraversal(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorderTraversal(node.Right)
	}

	inorderTraversal(root)
	return result
}

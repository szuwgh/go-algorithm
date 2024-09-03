package flatten

import "fmt"

// 二叉树展开为链表
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：

// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// 示例 1：

//      1
//     /  \
//    2    5
//   / \    \
//  3   4    6

//  1
//   \
//    2
//     \
//      3
// 	  \
//        4
//         \
//          5
// 		  \
// 		   6

// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
// 示例 2：

// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [0]
// 输出：[0]

// 要将二叉树展开为链表，可以使用前序遍历的思路，将树的每个节点处理成链表的一部分。具体步骤如下：

// 递归地展开左子树和右子树。
// 将左子树插入到右子树的地方，然后把原来的右子树接到新右子树的末端。
// 将左子树的指针置为null。
// 代码解释：
// flatten(root *TreeNode)：
// 首先递归地展开左子树和右子树。
// 然后将左子树插入到右子树的位置。
// 再将右子树连接到新右子树的末端。
// printList(root *TreeNode)：
// 用于打印展开后的链表结果。
// 示例：
// 对于输入的二叉树 [1,2,5,3,4,null,6]，
// 展开后的链表应该为 1 -> 2 -> 3 -> 4 -> 5 -> 6。
// 时间复杂度：
// 该算法的时间复杂度为 O(n)，其中 n 是二叉树中的节点个数。
// 每个节点最多被访问两次：一次是在递归展开时，另一次是在调整右子树位置时。
// 通过这种方式，你可以将二叉树按先序遍历顺序展开为链表
// TreeNode 定义二叉树节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flatten 函数将二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归展开左子树和右子树
	flatten(root.Left)
	flatten(root.Right)

	// 将左子树插入到右子树的位置
	tempRight := root.Right
	root.Right = root.Left
	root.Left = nil // 将左子树指针置为null

	// 将右子树接到新右子树的末端
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = tempRight
}

// 打印链表
func printList(root *TreeNode) {
	curr := root
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Right != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Right
	}
	fmt.Println()
}

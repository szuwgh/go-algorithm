package level_order

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// 示例 1：
//      3
//    /  \
//   9   20
//  / \  / \
//      15  7

// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 示例 2：

// 输入：root = [1]
// 输出：[[1]]
// 示例 3：

// 输入：root = []
// 输出：[]

// 代码说明
// TreeNode: 定义了二叉树的节点结构，包含值和左右子节点。
// levelOrder: 进行层序遍历的函数。
// 使用一个队列（queue）来存储当前层的节点。
// 遍历每一层，记录节点值，并将其子节点添加到队列中。
// 继续遍历直到队列为空。
// main: 创建一个示例树并调用 levelOrder 函数来执行层序遍历，最后打印结果。
// 通过这种方法，你可以实现二叉树的层序遍历并获取每一层的节点值。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 进行层序遍历并返回结果
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelValues []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 队列出队

			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValues)
	}

	return result
}

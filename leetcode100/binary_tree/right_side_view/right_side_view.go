package rightsideview

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例 1:

//      1
//     /  \
//    2    3
//   / \    \
//      5    4

// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
// 示例 2:

// 输入: [1,null,3]
// 输出: [1,3]

// 代码解析
// TreeNode 结构体: 定义二叉树的节点结构，
// 包含节点的值 Val，左子节点 Left 和右子节点 Right。
// rightSideView 函数: 返回从右侧看到的节点值列表。
// 使用 BFS 遍历树的每一层。
// 在每一层中，记录最后一个节点的值（即右视图中可见的节点）。
// main 函数: 构建示例二叉树，并调用 rightSideView 函数输出右视图节点值。
// 这个实现方法确保我们从树的顶部到底部逐层遍历，
// 并始终选择每层的最后一个节点作为右视图的一部分。

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView 返回二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 记录当前层的最后一个节点
			rightmost = node.Val

			// 将左子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// 将右子节点加入队列
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 将当前层的最后一个节点加入结果
		result = append(result, rightmost)
	}

	return result
}

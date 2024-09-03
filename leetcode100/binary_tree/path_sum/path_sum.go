package pathsum

//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// 示例 1：

//      10
//     /  \
//    5    -3
//   / \     \
//  3   2     11
// / \   \
// 3  -2  1

// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
// 示例 2：

// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3

// 这个问题可以用深度优先搜索（DFS）结合前缀和来解决。前缀和的思想是记录从根节点到当前节点的路径上，
// 所有节点值的和，并使用一个哈希表来存储路径和的出现次数。

// 具体步骤如下：

// 前缀和计算：我们定义prefixSum为从根节点到当前节点路径上所有节点值的和。
// 哈希表记录前缀和的出现次数：定义一个哈希表prefixSumCount，键为前缀和，值为该前缀和出现的次数。
// 路径和判断：对于每个节点，假设当前节点到根节点的路径和为currentSum，
// 如果currentSum - targetSum在哈希表中存在，则说明从某个祖先节点到当前节点的路径和等于targetSum。
// 递归遍历：递归地遍历每个节点的子树，并在回溯时更新哈希表中的前缀和。
// 代码详解：
// pathSum函数初始化了prefixSumCount并调用dfs进行深度优先搜索。
// dfs函数中，首先计算当前路径的和currentSum，然后根据哈希表查找是否有满足条件的前缀和，
// 接着递归遍历左右子树。在回溯的时候，更新哈希表以保证正确性。
// 最后在main函数中通过构建二叉树测试两个示例。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// 定义前缀和计数的哈希表
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1
	return dfs(root, 0, targetSum, prefixSumCount)
}

func dfs(node *TreeNode, currentSum int, targetSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	// 当前路径和
	currentSum += node.Val
	// 看看是否存在前缀和等于 currentSum - targetSum
	// 存在的话，说明从某个祖先节点到当前节点的路径和等于 targetSum
	numPathsToCurrent := prefixSumCount[currentSum-targetSum]

	// 更新路径和的哈希表
	prefixSumCount[currentSum]++

	// 递归遍历左右子树
	result := numPathsToCurrent + dfs(node.Left, currentSum, targetSum, prefixSumCount) + dfs(node.Right, currentSum, targetSum, prefixSumCount)

	// 回溯，去掉当前节点的路径和计数
	prefixSumCount[currentSum]--

	return result
}

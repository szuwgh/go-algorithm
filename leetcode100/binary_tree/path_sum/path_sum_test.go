package pathsum

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	// 示例1
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: -3}
	root1.Left.Left = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Left.Left = &TreeNode{Val: 3}
	root1.Left.Left.Right = &TreeNode{Val: -2}
	root1.Left.Right.Right = &TreeNode{Val: 1}
	root1.Right.Right = &TreeNode{Val: 11}
	fmt.Println(pathSum(root1, 8)) // 输出: 3

	// 示例2
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 4}
	root2.Right = &TreeNode{Val: 8}
	root2.Left.Left = &TreeNode{Val: 11}
	root2.Left.Left.Left = &TreeNode{Val: 7}
	root2.Left.Left.Right = &TreeNode{Val: 2}
	root2.Right.Left = &TreeNode{Val: 13}
	root2.Right.Right = &TreeNode{Val: 4}
	root2.Right.Right.Left = &TreeNode{Val: 5}
	root2.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pathSum(root2, 22)) // 输出: 3
}

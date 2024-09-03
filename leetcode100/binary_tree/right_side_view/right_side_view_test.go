package rightsideview

import (
	"fmt"
	"testing"
)

func TestRightsideview(t *testing.T) {
	// 示例 1
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}

	fmt.Println(rightSideView(root)) // 输出: [1, 3, 4]

	// 示例 2
	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(rightSideView(root)) // 输出: [1, 3]

	// 示例 3
	root = nil

	fmt.Println(rightSideView(root)) // 输出: []
}

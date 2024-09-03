package flatten

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	// 示例 1
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	flatten(root)
	printList(root) // 输出: 1 -> 2 -> 3 -> 4 -> 5 -> 6
}

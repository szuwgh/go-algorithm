package kthsmallest

import (
	"fmt"
	"testing"
)

func TestKthsmallest(t *testing.T) {
	// 创建示例二叉搜索树
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}

	k := 1
	fmt.Println(kthSmallest(root, k)) // 输出 1
}

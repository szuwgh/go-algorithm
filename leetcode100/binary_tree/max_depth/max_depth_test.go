package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	// 构建二叉树 [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("二叉树的最大深度:", maxDepth(root)) // 输出: 3
}

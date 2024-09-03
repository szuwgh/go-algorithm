package inorder_travesal

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root1 := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(inorderTraversal(root1)) // 输出: [1, 3, 2]

	// 示例 2
	var root2 *TreeNode = nil
	fmt.Println(inorderTraversal(root2)) // 输出: []

	// 示例 3
	root3 := &TreeNode{1, nil, nil}
	fmt.Println(inorderTraversal(root3)) // 输出: [1]
}

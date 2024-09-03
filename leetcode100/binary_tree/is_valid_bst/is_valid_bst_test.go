package is_valid_bst

import (
	"fmt"
	"testing"
)

func TestIsValidBst(t *testing.T) {
	// 示例用法
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	root2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}

	fmt.Println(isValidBST(root1)) // 输出: true
	fmt.Println(isValidBST(root2)) // 输出: false
}

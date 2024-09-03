package is_symmetric

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	// 示例1: root = [1,2,2,3,4,4,3]
	root1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}
	fmt.Println(isSymmetric(root1)) // 输出: true

	// 示例2: root = [1,2,2,null,3,null,3]
	root2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}
	fmt.Println(isSymmetric(root2)) // 输出: false
}

package level_order

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	// 创建一个示例树: [3,9,20,null,null,15,7]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// 执行层序遍历
	result := levelOrder(root)
	fmt.Println(result) // 输出: [[3], [9, 20], [15, 7]]
}

package buildtree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	fmt.Print("前序遍历结果: ")
	printTree(root) // 输出：3 9 20 15 7
}

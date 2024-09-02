package tree

import (
	"fmt"
	"testing"
)

/*

			      3
			 /        \
		    9         20
		  /	 \	     /   \
		 8	  6	    15    7
	   /  \  / \   / \   / \
	  10  11 12 13 14 16 17 18
*/
func Test_BinarySearchTree(t *testing.T) {
	n := &TreeNode{3, nil, nil}
	n.Left = &TreeNode{9, nil, nil}
	n.Right = &TreeNode{20, nil, nil}
	n.Left.Left = &TreeNode{8, nil, nil}
	n.Left.Left.Left = &TreeNode{10, nil, nil}
	n.Left.Left.Right = &TreeNode{11, nil, nil}

	n.Left.Right = &TreeNode{6, nil, nil}
	n.Left.Right.Left = &TreeNode{12, nil, nil}
	n.Left.Right.Right = &TreeNode{13, nil, nil}

	n.Right.Left = &TreeNode{15, nil, nil}
	n.Right.Left.Left = &TreeNode{14, nil, nil}
	n.Right.Left.Right = &TreeNode{16, nil, nil}

	n.Right.Right = &TreeNode{7, nil, nil}
	n.Right.Right.Left = &TreeNode{17, nil, nil}
	n.Right.Right.Right = &TreeNode{18, nil, nil}

	fmt.Print("前序遍历： ")
	n.PreOrder()
	fmt.Print("\n")
	fmt.Print("中序遍历： ")
	n.MiddleOrder()
	fmt.Print("\n")
	fmt.Print("后序遍历： ")
	n.PostOrder()
	fmt.Print("\n")
	fmt.Print("蛇形遍历： \n")
	n.zigzagLevelOrder()
}

func Test_arr(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = append(arr[:1], arr...)
	fmt.Print(arr)
}

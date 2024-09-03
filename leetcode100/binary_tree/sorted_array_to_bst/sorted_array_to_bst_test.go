package sorted_array_to_bst

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums1 := []int{-10, -3, 0, 5, 9}
	nums2 := []int{1, 3}

	tree1 := sortedArrayToBST(nums1)
	tree2 := sortedArrayToBST(nums2)

	fmt.Println("Tree 1:")
	printTree(tree1)

	fmt.Println("Tree 2:")
	printTree(tree2)
}

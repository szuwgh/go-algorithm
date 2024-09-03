package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// 示例 1: [1,2,2,1]
	head1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(isPalindrome(head1)) // 输出: true

	// 示例 2: [1,2]
	head2 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(isPalindrome(head2)) // 输出: false
}

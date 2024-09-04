package long

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad")) // 输出 "bab" 或 "aba"
	fmt.Println(longestPalindrome("cbbd"))  // 输出 "bb"
}

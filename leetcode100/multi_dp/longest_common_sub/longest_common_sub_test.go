package long

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2)) // 输出 3

	text1 = "abc"
	text2 = "abc"
	fmt.Println(longestCommonSubsequence(text1, text2)) // 输出 3

	text1 = "abc"
	text2 = "def"
	fmt.Println(longestCommonSubsequence(text1, text2)) // 输出 0
}

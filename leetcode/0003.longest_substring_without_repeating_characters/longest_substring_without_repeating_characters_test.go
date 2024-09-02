package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {

	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
}

package wordbreak

import (
	"fmt"
	"testing"
)

func TestWorkBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                       // true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))                  // true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

package longest_palindromic_substring

import (
	"fmt"
	"testing"
)

func Test_longestPalinddrome(t *testing.T) {
	s := "cabad"
	fmt.Println(s, longestPalinddrome(s))
	s = "cbbd"
	fmt.Println(s, longestPalinddrome(s))

}

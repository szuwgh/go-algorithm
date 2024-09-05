package isValid

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
}

package rob

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	n := rob([]int{1, 2, 3, 1})
	fmt.Println(n)

	n = rob([]int{2, 7, 9, 3, 1})
	fmt.Println(n)
}

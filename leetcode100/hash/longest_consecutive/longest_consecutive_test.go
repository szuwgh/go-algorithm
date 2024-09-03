package longestconsecutive

import (
	"fmt"
	"testing"
)

func Test_longest_conse(t *testing.T) {
	n := longestconsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(n)

	n = longestconsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Println(n)

}

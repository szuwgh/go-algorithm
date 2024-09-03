package movezero

import (
	"fmt"
	"testing"
)

func Test_movezero(t *testing.T) {
	num := []int{0, 1, 0, 3, 12}
	movezero(num)
	fmt.Println(num)

	num = []int{5, 1, 2, 3, 12}
	movezero(num)
	fmt.Println(num)
}

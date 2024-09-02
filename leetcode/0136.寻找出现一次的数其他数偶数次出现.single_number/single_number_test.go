package single_number

import (
	"fmt"
	"testing"
)

func Test_SingleNumber(t *testing.T) {
	a := []int{4, 1, 2, 1, 2}
	x := singleNumber(a)
	fmt.Println(x)
}

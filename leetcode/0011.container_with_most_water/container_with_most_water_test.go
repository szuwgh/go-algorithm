package container_with_most_water

import (
	"fmt"
	"testing"
)

func Test_maxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("max area is:", maxArea(height))
}

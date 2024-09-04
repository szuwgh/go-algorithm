package min

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2)) // 输出: 3

	word1 = "intention"
	word2 = "execution"
	fmt.Println(minDistance(word1, word2)) // 输出: 5
}

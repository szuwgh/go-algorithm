package coin_change

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount)) // 输出: 3
}

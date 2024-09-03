package max_profit

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println("Maximum profit for prices1:", maxProfit(prices1)) // 输出: 5
	fmt.Println("Maximum profit for prices2:", maxProfit(prices2)) // 输出: 0
}

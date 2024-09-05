package daily_temp

import (
	"fmt"
	"testing"
)

func TestDailyTemp(t *testing.T) {
	temperatures1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures2 := []int{30, 40, 50, 60}
	temperatures3 := []int{30, 60, 90}

	fmt.Println(dailyTemperatures(temperatures1)) // 输出: [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures(temperatures2)) // 输出: [1 1 1 0]
	fmt.Println(dailyTemperatures(temperatures3)) // 输出: [1 1 0]
}

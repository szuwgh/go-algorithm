package container_with_most_water

// 题目描述：

// 给定 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
// 找出其中的两条线，
// 使得它们与 x 轴共同构成的容器可以容纳最多的水。
// **说明：**你不能倾斜容器，且 n 的值至少为 2。

/*



  8_|
  7_|       |___________________|________
  6_|       |                   |       |
  5_|       |   |               |       |
  4_|       |   |       |       |       |
  3_|       |   |       |   |   |       |
  2_|       |   |       |   |   |   |   |
  1_|       |   |   |   |   |   |   |   |
  0_|___|___|___|___|___|___|___|___|___|___

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)
		area := h * (j - i)
		if area > max {
			max = area
		}
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

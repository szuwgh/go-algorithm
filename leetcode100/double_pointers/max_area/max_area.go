package main

import (
	"math"
)

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 返回容器可以储存的最大水量。

// 说明：你不能倾斜容器。

// 示例 1：

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：

// 输入：height = [1,1]
// 输出：1

// 解题思路：
// 初始化两个指针，分别指向数组的左右两端。
// 计算两个指针对应的高度和距离所能容纳的水量，然后更新最大水量。
// 将高度较小的指针向内移动一步（因为水量由较小的高度决定，移动较小高度的一侧可能会找到更大的容器）。
// 重复上述过程，直到两个指针相遇。

// left 和 right 分别是指向数组最左端和最右端的指针。
// 在每一步中，计算以 left 和 right 为边界的容器可以容纳的水量，并尝试更新最大水量。
// 根据高度，决定是移动左指针还是右指针，目的是尝试找到更大的容器。
// 最后，当 left 与 right 相遇时，循环结束，返回 maxWater，即最大水量。

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// 计算当前的水量
		h := int(math.Min(float64(height[left]), float64(height[right])))
		water := h * (right - left)
		// 更新最大水量
		if water > maxWater {
			maxWater = water
		}

		// 移动较低的一侧的指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

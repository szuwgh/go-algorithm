package median_of_two_sorted_arrays

// 4.寻找两个有序数组的中位数
// 题目描述：
// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 则中位数是 2.0
// 1
// 2
// 3
// 4

// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]

// 则中位数是 (2 + 3)/2 = 2.5

func medianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	result := mergeSort(nums1, nums2)
	return medianNum(result)
}

func mergeSort(nums1 []int, nums2 []int) []int {
	result := []int{}
	m, n := 0, 0
	l, r := len(nums1), len(nums2)
	for m < l && n < r {
		if nums1[m] < nums2[n] {
			result = append(result, nums1[m])
			m++
		} else {
			result = append(result, nums2[n])
			n++
		}
	}
	result = append(result, nums2[n:]...)
	result = append(result, nums1[m:]...)
	return result
}

//中位数
func medianNum(num []int) float64 {
	l := len(num)
	if l == 0 {
		return 0
	}
	if l%2 == 0 {
		return float64(num[l/2]+num[l/2-1]) / 2.0
	}
	return float64(num[l/2])
}

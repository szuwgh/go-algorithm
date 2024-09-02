package sort

import "fmt"

var arr = []int{8, 4, 2, 9, 10, 3, 5, 20, 15, 1}

//冒泡
//8 4 2 9 10 3 5 20 15 1
//4 2 8 9 3 5 10 15 1 20
//2 4 8 3 5 9 10 1 15 20
//2 4 3 5 8 9 1 10 15 20
//2 3 4 5 8 1 9 10 15 20
//2 3 4 5 1 8 9 10 15 20
//2 3 4 1 5 8 9 10 15 20
//2 3 1 4 5 8 9 10 15 20
//2 1 3 4 5 8 9 10 15 20
//1 2 3 4 5 8 9 10 15 20
func MaoPaoSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}
}

//快速排序 
// [8 4 2 9 10 3 5 20 15 1]
// [3 4 2 1 5 8 10 20 15 9]
// [2 1 3 4 5 8 10 20 15 9]
// [1 2 3 4 5 8 10 20 15 9]
// [1 2 3 4 5 8 10 20 15 9]
// [1 2 3 4 5 8 9 10 15 20]
// [1 2 3 4 5 8 9 10 15 20]
// [1 2 3 4 5 8 9 10 15 20]
func QuickSort(arr []int) {
	quick(arr, 0, len(arr)-1)

}

func quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	key := arr[left] //选择第一个数为基准
	i, j := left, right
	for i < j {
		//从右向左找小于key的数
		for i < j && arr[j] >= key { //j=9,6,5
			j--
		}
		//从左向右找大于key的数
		for i < j && arr[i] <= key { //i=3,4,5
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	// i=j结束扫描
	// 基准数归位，现在左边的序列小于基准数，右边的序列大于基准数
	arr[left], arr[i] = arr[i], arr[left]
	fmt.Println(arr)
	quick(arr, left, i-1)  //继续排左边
	quick(arr, i+1, right) //继续排右边
}

//选择排序
// [8 4 2 9 10 3 5 20 15 1]
// [1 4 2 9 10 3 5 20 15 8]
// [1 2 4 9 10 3 5 20 15 8]
// [1 2 3 9 10 4 5 20 15 8]
// [1 2 3 4 10 9 5 20 15 8]
// [1 2 3 4 5 9 10 20 15 8]
// [1 2 3 4 5 8 10 20 15 9]
// [1 2 3 4 5 8 9 20 15 10]
// [1 2 3 4 5 8 9 10 15 20]
// [1 2 3 4 5 8 9 10 15 20]
// [1 2 3 4 5 8 9 10 15 20]
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
		fmt.Println(arr)
	}
}

// [8 4 2 9 10 3 5 20 15 1]
// [4 8 2 9 10 3 5 20 15 1]
// [2 4 8 9 10 3 5 20 15 1]
// [2 4 8 9 10 3 5 20 15 1]
// [2 4 8 9 10 3 5 20 15 1]
// [2 3 4 8 9 10 5 20 15 1]
// [2 3 4 5 8 9 10 20 15 1]
// [2 3 4 5 8 9 10 20 15 1]
// [2 3 4 5 8 9 10 15 20 1]
// [1 2 3 4 5 8 9 10 15 20]
// [1 2 3 4 5 8 9 10 15 20]
func InsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		fmt.Println(arr)
	}
}

// [8 4 2 9 10 3 5 20 15 1]
// [4 8]
// [9 10]
// [2 9 10]
// [2 4 8 9 10]
// [3 5]
// [1 15]
// [1 15 20]
// [1 3 5 15 20]
// [1 2 3 4 5 8 9 10 15 20]
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	fmt.Println(result)
	return result
}

func merge(left, right []int) []int {
	var result []int
	m, n := 0, 0 // left和right的index位置
	//两个有序数组合并排序方法
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

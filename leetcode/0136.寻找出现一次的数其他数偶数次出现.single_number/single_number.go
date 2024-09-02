package single_number

func singleNumber(num []int) int {
	result := 0
	for _, v := range num {
		result ^= v
	}
	return result
}

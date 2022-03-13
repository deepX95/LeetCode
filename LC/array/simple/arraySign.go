package simple

// arraySign 1822. 数组元素积的符号
func arraySign(nums []int) int {
	ret := 1
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] > 0 {
			ret *= 1
		} else {
			ret *= -1
		}
	}
	return ret
}

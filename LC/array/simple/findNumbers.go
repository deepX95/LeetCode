package simple

// findNumbers 1295. 统计位数为偶数的数字
func findNumbers(nums []int) int {
	cnt := 0
	for i := range nums {
		cur := 0
		v := nums[i]
		for v > 0 {
			v = v / 10
			cur++
		}
		if cur%2 == 0 {
			cnt++
		}
	}
	return cnt

}

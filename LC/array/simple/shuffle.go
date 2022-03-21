package simple

// shuffle 1470. 重新排列数组
func shuffle(nums []int, n int) []int {
	l := len(nums)
	ret := make([]int, l)
	for i := 0; i < n; i++ {
		ret[2*i] = nums[i]
		ret[2*i+1] = nums[i+n]
	}
	return ret
}

package simple

// maximumDifference 2016. 增量元素之间的最大差值
func maximumDifference(nums []int) int {
	max, min := -1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > min {
			max = maxV(max, nums[i]-min)
		} else {
			min = nums[i]
		}
	}
	return max
}

func maxV(a, b int) int {
	if a > b {
		return a
	}
	return b
}

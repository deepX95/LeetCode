package simple

// maxSubArray 面试题 16.17. 连续数列
// fixme:分治算法
func maxSubArray(nums []int) int {
	// dp
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxMaxSubArray(a, b int) int {
	if a > b {
		return a
	}
	return b
}

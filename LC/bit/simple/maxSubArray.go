package simple

// maxSubArray 面试题 16.17. 连续数列
// fixme:分治算法
func maxSubArray(nums []int) int {
	// dp
	//max := nums[0]
	//for i := 1; i < len(nums); i++ {
	//	if nums[i]+nums[i-1] > nums[i] {
	//		nums[i] += nums[i-1]
	//	}
	//	if nums[i] > max {
	//		max = nums[i]
	//	}
	//}
	//return max

	// 贪心算法
	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i])
		maxSum = max(curSum, maxSum)
	}
	return maxSum

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

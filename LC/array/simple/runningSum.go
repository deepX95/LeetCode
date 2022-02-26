package simple

// runningSum 1480. 一维数组的动态和
func runningSum(nums []int) []int {
	for i:=1;i<len(nums);i++{
		nums[i]+=nums[i-1]
	}
	return nums
}
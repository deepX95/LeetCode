package simple

import "math"

// MaxSubArray 53. 最大子数组和
// fixme:https://leetcode-cn.com/problems/maximum-subarray/solution/dong-tai-gui-hua-fen-zhi-fa-python-dai-ma-java-dai/
func maxSubArray(nums []int) int {
	//dp:=make([]int,len(nums))
	//dp[0]=nums[0]
	//n:=len(nums)
	//for i:=1;i<n;i++{
	//	if dp[i-1]>=0{
	//		dp[i]=dp[i-1]+nums[i]
	//	}else{
	//		dp[i]=nums[i]
	//	}
	//}
	//return maxSlice(dp)

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxSlice(a []int) int {
	max:=math.MinInt16
	for i:=range a{
		if a[i]>max{
			max=a[i]
		}
	}
	return max
}

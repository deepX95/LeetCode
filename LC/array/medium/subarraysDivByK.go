package medium

// subarraysDivByK 974. 和可被 K 整除的子数组
// fixme:多看几次
func subarraysDivByK(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		preSum := 0
		for j := i; j < len(nums); j++ {
			preSum += nums[j]
			if preSum%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

package simple

// smallerNumbersThanCurrent 1365. 有多少小于当前数字的数字
// fixme:计数排序
func smallerNumbersThanCurrent(nums []int) []int {
	cnt := [101]int{}
	// 数字 v 出现的次数
	for _, v := range nums {
		cnt[v]++
	}
	// 统计值<=i在数组中出现的个数
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}
	ans := make([]int, len(nums))
	// 将值<nums[i]的数字个数及cntk-1的值放入最后的结果数组
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}

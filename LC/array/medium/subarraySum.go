package medium

// subarraySum 560. 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	// 模拟
	//cnt := 0
	//for start := 0; start < len(nums); start++ {
	//	sum:=0
	//	for end:=start;end>=0;end-- {
	//		sum += nums[end]
	//		if sum == k {
	//			cnt++
	//		}
	//	}
	//}
	//return cnt

	// 前缀和
	preSum,res:=0,0
	visited:=map[int]int{
		0:1,
	}
	for i:=0;i<len(nums);i++{
		preSum+=nums[i]
		if v,ok:=visited[preSum-k];ok{
			res+=v
		}
		visited[preSum]+=1
	}
	return res
}

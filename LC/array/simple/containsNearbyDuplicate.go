package simple

// containsNearbyDuplicate 219. 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	// 暴力不通过
	//for i:=0;i<len(nums);i++{
	//	for j:=i+1;j<len(nums);j++{
	//		if nums[i]==nums[j] && j-i<=k{
	//			return true
	//		}
	//	}
	//}
	//return false

	// 哈希
	mp := map[int]int{}
	for i, num := range nums {
		if p, ok := mp[num]; ok && i-p <= k {
			return true
		}
		mp[num] = i
	}
	return false

}

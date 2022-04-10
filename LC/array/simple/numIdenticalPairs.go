package simple

// numIdenticalPairs 1512. 好数对的数目
func numIdenticalPairs(nums []int) (cnt int) {
	//for i:=0;i<len(nums);i++{
	//	for j:=i+1;j<len(nums);j++{
	//		if nums[i]==nums[j]{
	//			cnt++
	//		}
	//	}
	//}
	//return cnt

	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				cnt++
			}
		}
	}
	return

}

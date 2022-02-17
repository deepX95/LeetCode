package simple

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	mp:=make(map[int]int,len(nums))
	for i:=range nums{
		if v,ok:=mp[target -nums[i]];ok{
			return []int{i,v}
		}
		mp[nums[i]]=i
	}

	return nil
}
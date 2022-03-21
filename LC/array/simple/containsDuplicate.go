package simple

// containsDuplicate 217. 存在重复元素
func containsDuplicate(nums []int) bool {
	mp := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			return true
		}
		mp[v] = struct{}{}
	}
	return false
}

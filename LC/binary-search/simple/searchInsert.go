package simple

// searchInsert 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	//l, r := 0, len(nums)-1
	//for l < r {
	//	mid := l + (r-l)/2
	//	if nums[mid] < target {
	//		l = mid+1
	//	} else {
	//		r = mid
	//	}
	//}
	//return l

	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid+1
		}
	}
	return l
}

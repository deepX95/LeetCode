package medium

// searchRange 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPos := findFirstPosition(nums, target)
	if firstPos == -1 {
		return []int{-1, -1}
	}
	lastPos := findLastPosition(nums, target)

	return []int{firstPos, lastPos}
}

// 查找第一个元素的位置
func findFirstPosition(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			// nums[mid]>=target
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func findLastPosition(nums []int, target int) int {
	l,r,mid:=0,len(nums)-1,0
	for l<r{
		mid=l+(r-l+1)/2
		if nums[mid]>target{
			// [left..mid - 1]
			r=mid-1
		}else{
			// 下一轮搜索区间是 [mid..right]
			// nums[mid]<=target
			l=mid
		}
	}

	return l
}

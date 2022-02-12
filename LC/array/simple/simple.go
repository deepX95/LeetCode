package simple

// MoveZeroes 283. 移动零
// fixme:再做一次
func MoveZeroes(nums []int)  {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}

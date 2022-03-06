package medium

// rotate 189. 轮转数组
func rotate(nums []int, k int) {
	// 1.暴力
	//newNums := make([]int, len(nums))
	//n := len(nums)
	//for i := range nums {
	//	newNums[(i+k)%n] = nums[i]
	//}
	//copy(nums, newNums)

	// 2.数组翻转
	k%=len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

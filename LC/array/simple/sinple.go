package simple

// Search 704. 二分查找
// fixme:https://leetcode-cn.com/problems/binary-search/solution/er-fen-cha-zhao-xiang-jie-by-labuladong/
// https://www.zhihu.com/question/36132386/answer/530313852
func Search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l <= r {
		mid := l + r>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r--
		} else {
			l++
		}
	}
	return -1
}

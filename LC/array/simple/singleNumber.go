package simple

// SingleNumber 136. 只出现一次的数字
// fixme:位运算，0异或任何数都是本身，任何数异或自己都是0
func SingleNumber(nums []int) int {
	single:=0
	for i:=range nums{
		single^=nums[i]
	}
	return single
}
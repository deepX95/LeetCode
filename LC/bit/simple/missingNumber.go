package simple

// missingNumber 268. 丢失的数字
// fixme:二分查找
func missingNumber(nums []int) (xor int) {
	// 位运算
	//for i, num := range nums {
	//	xor ^= i ^ num
	//}
	//return xor ^ len(nums)

	// 数学
	n:=len(nums)
	sum:=n*(n+1)/2
	for i:=range nums{
		sum-=nums[i]
	}
	return sum

	// 二分查找

}

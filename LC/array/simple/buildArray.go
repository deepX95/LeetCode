package simple

// buildArray 1920. 基于排列构建数组
// fixme:原地修改
func buildArray(nums []int) []int {
	//n:=len(nums)
	//ret:=make([]int,n)
	//for i:=range nums{
	//	ret[i]=nums[nums[i]]
	//}
	//return ret

	// 原地修改
	n := len(nums)
	// 第一次遍历编码最终值
	for i := 0; i < n; i++ {
		nums[i] += 1000 * (nums[nums[i]] % 1000)
	}
	// 第二次遍历修改为最终值
	for i := 0; i < n; i++ {
		nums[i] /= 1000
	}
	return nums
}

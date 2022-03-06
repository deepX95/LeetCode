package medium

// singleNumber 137. 只出现一次的数字 II
// fixme：位运算 / 状态机
func singleNumber(nums []int) int {
	// 1.暴力
	//freq := map[int]int{}
	//for _, v := range nums {
	//	freq[v]++
	//}
	//for num, occ := range freq {
	//	if occ == 1 {
	//		return num
	//	}
	//}
	//return 0

	// 位运算
	// fixme:https://leetcode-cn.com/problems/single-number-ii/solution/single-number-ii-mo-ni-san-jin-zhi-fa-by-jin407891/
	// 负数会错误
	counts := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			counts[j] += nums[i] & 1 // 更新第 j 位
			nums[i] >>= 1 //  第 j 位 --> 第 j + 1 位
		}
	}
	res,m:=0,3
	for i:=0;i<32;i++{
		res<<=1
		res |=counts[31-i]%m
	}
	return res

}

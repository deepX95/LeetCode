package simple

import "sort"

// maxSubsequence 2099. 找到和最大的长度为 K 的子序列
func maxSubsequence(nums []int, k int) []int {
	n:=len(nums)
	tmp:=make([][]int,n)
	for i,num:=range nums{
		tmp[i]=[]int{i,num}
	}
	sort.Slice(tmp,func(i,j int)bool{
		return tmp[i][1]>tmp[j][1]
	})

	b := make([][]int, k)
	for i := 0; i < k; i++ {
		b[i] = []int{tmp[i][0], tmp[i][1]}
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = b[i][1]
	}
	return res
}

package simple

// findDisappearedNumbers 448. 找到所有数组中消失的数字
// fixme:看很多新奇的题解
func findDisappearedNumbers(nums []int) (ret []int) {
	// 法 1
	//n:=len(nums)
	//mp:=make(map[int]struct{},n)
	//for i:=range nums{
	//	mp[nums[i]]= struct{}{}
	//}
	//for i:=1;i<=n;i++{
	//	if _,ok:=mp[i];!ok{
	//		ret=append(ret,i)
	//	}
	//}
	//return

	// 法2
	// fixme:再研究下
	n:=len(nums)
	for _,v:=range nums{
		v=(v-1)%n
		nums[v]+=n
	}
	for i,v:=range nums{
		if v<=n{
			ret=append(ret,i+1)
		}
	}
	return
}
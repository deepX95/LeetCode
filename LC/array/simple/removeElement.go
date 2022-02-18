package simple

// removeElement 27. 移除元素
func removeElement(nums []int, val int) int {
	// 循环不变量：nums[0...j)!=val
	// j 指向了下一个要赋值的元素的位置
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[j]=nums[i]
			j++
		}
	}
	return j

	// 循环不变量：nums[0...j]!=val
	// j 指向了上一个已赋值的元素的位置
	//j:=-1
	//for i:=0;i<len(nums);i++{
	//	if nums[i]!=val{
	//		j++
	//		nums[j]=nums[i]
	//	}
	//
	//}
	//return j+1
}
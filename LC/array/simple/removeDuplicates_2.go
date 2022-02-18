package simple

// removeDuplicates_2 80. 删除有序数组中的重复项 II
func removeDuplicates_2(nums []int) int {
	n:=len(nums)
	if n<2{
		return  n
	}

	// 循环不变量 nums[0...j)是有序的，并且相同元素最多保留2次
	// j 指向下一个要赋值的元素的位置
	j:=2
	for i:=2;i<n;i++{
		if nums[i]!=nums[j-2]{
			nums[j]=nums[i]
			j++
		}
	}
	return j

	//n:=len(nums)
	//if n<2{
	//	return  n
	//}
	//
	//// 循环不变量 nums[0...j]是有序的，并且相同元素最多保留2次
	//// j 已经赋值过的元素的最后一个位置
	//j:=1
	//for i:=2;i<n;i++{
	//	if nums[i]!=nums[j-1]{
	//		j++
	//		nums[j]=nums[i]
	//	}
	//}
	//return j+1
}
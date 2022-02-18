package simple

//removeDuplicates 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	// nums[0..j]有序，且元素唯一
	// j 表示上一个赋值的元素的下标
	//j:=0
	//for i:=1;i<len(nums);i++{
	//	if nums[i]!=nums[j]{
	//		j++
	//		nums[j]=nums[i]
	//	}
	//}
	//return j+1

	// nums[0..j)有序，且元素唯一
	//j 表示下一个需要赋值的元素的下标
	j:=1
	for i:=1;i<len(nums);i++{
		if nums[i]!=nums[j-1]{
			nums[j]=nums[i]
			j++
		}
	}
	return j

}
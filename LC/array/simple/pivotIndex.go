package simple

// pivotIndex 724. 寻找数组的中心下标
func pivotIndex(nums []int) int {
	total:=0
	for _,v:=range nums{
		total+=v
	}
	sum:=0
	for i,v:=range nums{
		if 2*sum+v==total{
			return i
		}
		sum+=v
	}
	return -1
}
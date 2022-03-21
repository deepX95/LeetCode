package simple

type NumArray struct {
	PreSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}
	return NumArray{PreSum: preSum}
}

// SumRange 303. 区域和检索 - 数组不可变
func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}

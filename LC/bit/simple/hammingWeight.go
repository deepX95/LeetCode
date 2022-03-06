package simple

// hammingWeight 191. 位1的个数
func hammingWeight(num uint32) (cnt int) {
	// 循环检查二进制位
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			cnt++
		}
	}
	return

	// 位运算优化
	//for ; num > 0; num &= num - 1 {
	//	cnt++
	//}
	//return
}

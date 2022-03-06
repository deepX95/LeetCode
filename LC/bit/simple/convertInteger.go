package simple

// convertInteger 面试题 05.06. 整数转换
func convertInteger(A int, B int) int {
	// 0000 0001
	// 0101 0010
	// 0000 0011

	// 0001 1101
	// 0000 1111

	// 0001 0010   0
	// 0000 1001   1
	// 0000 0100   2
	// 0000 0010   3
	// 0000 0001   4

	// 位运算
	xor:=A^B
	cnt:=0
	for i:=0;i<32;i++{
		if xor&1>0{
			cnt++
		}
		xor>>=1
	}
	return cnt
}
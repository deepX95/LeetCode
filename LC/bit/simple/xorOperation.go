package simple

// xorOperation 1486. 数组异或操作
// fixme:数学推导
func xorOperation(n int, start int) int {
	ret := 0
	for i := 0; i < n; i++ {
		ret ^= start + 2*i
	}
	return ret
}

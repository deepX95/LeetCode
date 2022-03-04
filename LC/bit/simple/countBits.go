package simple

// countBits 338. 比特位计数
// fixme:dp
func countBits(n int) []int {
	// 位运算
	ret := make([]int, n+1)
	for i := range ret {
		cnt := 0
		for j := i; j > 0; j &= j - 1 {
			cnt++
		}
		ret[i]= cnt
	}
	return ret
}

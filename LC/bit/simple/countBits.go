package simple

// countBits 338. 比特位计数
// fixme:dp
func countBits(n int) []int {
	// 位运算
	ret := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		cnt := 0
		for j := i; j > 0; j &= j - 1 {
			cnt++
		}
		ret = append(ret, cnt)
	}
	return ret
}

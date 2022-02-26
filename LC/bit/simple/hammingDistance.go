package simple

// hammingDistance  461. 汉明距离
func hammingDistance(x, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return

	//for s := x ^ y; s > 0; s &= s - 1 {
	//	ans++
	//}
	//return

	// 库函数
	//return bits.OnesCount(uint(x ^ y))
}

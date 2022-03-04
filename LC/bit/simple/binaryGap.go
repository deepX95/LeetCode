package simple

// binaryGap 868. 二进制间距
func binaryGap(n int) int {
	ans, last := 0, -1
	for i := 0; i < 32; i++ {
		if (n >> i & 1) > 0 {
			if last >= 0 {
				ans = maxBinaryGap(ans, i-last)
			}
			last = i
		}
	}
	return ans
}

func maxBinaryGap(a, b int) int {
	if a > b {
		return a
	}
	return b
}

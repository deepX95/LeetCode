package simple

// divisorGame 1025. 除数博弈
func divisorGame(n int) bool {
	f := make([]bool, n+5)
	f[1], f[2] = false, true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}

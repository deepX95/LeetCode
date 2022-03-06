package simple

// isPowerOfFour 342. 4的幂
func isPowerOfFour(n int) bool {
	// 2的幂次方，4的幂取模一定是1
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
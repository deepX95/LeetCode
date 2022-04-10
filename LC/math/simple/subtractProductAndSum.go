package simple

// subtractProductAndSum  1281. 整数的各位积和之差
func subtractProductAndSum(n int) int {
	mul, add := 1, 0
	for n != 0 {
		mul *= n % 10
		add += n % 10
		n /= 10
	}
	return mul - add
}

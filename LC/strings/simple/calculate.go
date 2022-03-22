package simple

// calculate LCP 17. 速算机器人
func calculate(s string) int {
	// x:1 y:0
	//"A" 运算：使 x = 2 * x + y；
	//"B" 运算：使 y = 2 * y + x。
	x, y := 1, 0
	for i := range s {
		switch s[i] {
		case 'A':
			x = 2*x + y
		case 'B':
			y = 2*y + x
		}
	}
	return x + y
}

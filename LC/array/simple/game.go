package simple

// game  LCP 01. 猜数字
func game(guess []int, answer []int) int {
	r := 0
	for i := range guess {
		if guess[i] == answer[i] {
			r++
		}
	}
	return r
}

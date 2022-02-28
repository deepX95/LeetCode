package simple

// IsPalindrome 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverNum := 0
	num:=x
	for num != 0 {
		reverNum = reverNum*10 + num%10
		num /= 10
	}
	if x == reverNum{
		return true
	}
	return false
}
package simple

// firstPalindrome 2108. 找出数组中的第一个回文字符串
func firstPalindrome(words []string) string {
	for i:=range words{
		if isPalindrome(words[i]){
			return words[i]
		}
	}
	return ""
}

func isPalindrome(word string)bool{
	l,r:=0,len(word)
	for l<r{
		if word[l]!=word[r]{
			return false
		}
		l++
		r--
	}
	return true
}

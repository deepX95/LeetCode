package simple

import "unicode"

// detectCapitalUse 520. 检测大写字母
func detectCapitalUse(word string) bool {
	//if unicode.IsLower(rune(word[0])) && unicode.IsLower(rune(word[1])){
	//	return false
	//}
	//
	//for i:=2;i<len(word);i++{
	//	if unicode.IsLower(rune(word[1]))!=unicode.IsLower(rune(word[i])){
	//		return false
	//	}
	//}
	//return true
	cnt, n := 0, len(word)
	for i := 0; i < n; i++ {
		if unicode.IsUpper(rune(word[i])) {
			cnt++
		}
	}
	return cnt == n || cnt == 0 || (cnt == 1 && unicode.IsUpper(rune(word[0])))
}

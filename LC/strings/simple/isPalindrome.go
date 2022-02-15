package simple

import (
	"strings"
)

// IsPalindrome 125. 验证回文串
// fixme:unicode使用，空白字符串
// fixme:复看
func IsPalindrome(s string) bool {
	// 双指针
	//s = strings.ToLower(s)
	//l,r:=0,len(s)-1
	//for l<r{
	//	//for l<r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
	//	//	l++
	//	//}
	//	//
	//	//for l<r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
	//	//	r--
	//	//}
	//
	//	for l < r && !isalnum(s[l]) {
	//		l++
	//	}
	//	for l < r && !isalnum(s[r]) {
	//		r--
	//	}
	//
	//
	//	if l<r {
	//		ls,rs:=s[l],s[r]
	//		ll,rr:=string(ls),string(rs)
	//		fmt.Println(ll,rr)
	//		if ls !=rs {
	//			return false
	//		}
	//		l++
	//		r--
	//	}
	//}
	//return true

	// 双指针
	//if len(s) == 0 {
	//	return true
	//}
	//i,j := 0,len(s)-1
	//
	//for j > i {
	//	for j > i && !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
	//		i++
	//	}
	//	for j > i && !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
	//		j--
	//	}
	//	if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
	//		return false
	//	}
	//	i++
	//	j--
	//}
	//return true

	// 筛选 + 判断
	var sblder strings.Builder
	for i:=0;i<len(s);i++{
		if isalnum(s[i]){
			sblder.WriteByte(s[i])
		}
	}
	sgood:=sblder.String()

	n:=len(sgood)
	sgood=strings.ToLower(sgood)
	for i:=0;i<n/2;i++{
		if sgood[i]!=sgood[n-1-i]{
			return false
		}
	}
	return true

}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
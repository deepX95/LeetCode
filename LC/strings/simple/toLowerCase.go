package simple

import (
	"strings"
)

// toLowerCase 709. 转换成小写字母
func toLowerCase(s string) string {
	// 模拟
	//res := stringToBytes(s)
	//for i := range res {
	//	if 'A' <= res[i] && res[i] <= 'Z' {
	//		//res[i] += 'a' - 'A'
	//		res[i] += 32
	//	}
	//}
	//return *(*string)(unsafe.Pointer(&res))

	// 位运算
	lower := &strings.Builder{}
	lower.Grow(len(s))
	for _, ch := range s {
		if 65 <= ch && ch <= 90 {
			ch |= 32
		}
		lower.WriteRune(ch)
	}
	return lower.String()
}

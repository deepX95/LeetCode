package simple

import (
	"strings"
	"unsafe"
)

// reversePrefix 2000. 反转单词前缀
func reversePrefix(word string, ch byte) string {
	tgt := strings.IndexByte(word, ch)
	res := stringToBytesReversePrefix(word)
	l := 0
	for l < tgt {
		res[l], res[tgt] = res[tgt], res[l]
		l++
		tgt--
	}
	return *(*string)(unsafe.Pointer(&res))
}

func stringToBytesReversePrefix(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

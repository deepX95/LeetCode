package simple

import "unsafe"

// replaceDigits 1844. 将所有数字用字符替换
func replaceDigits(s string) string {
	res := stringToBytes(s)
	for i := 1; i < len(res); i += 2 {
		res[i] = res[i-1] + res[i] - '0'
	}
	return *(*string)(unsafe.Pointer(&res))
}

func stringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

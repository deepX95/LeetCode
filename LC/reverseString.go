package LC

import "fmt"

// ReverseString 反转字符串
func ReverseString(str []string) {
	for left, right := 0, len(str)-1; left < right; left++ {
		str[left], str[right] = str[right], str[left]
		right--
	}
	fmt.Println(str)
}

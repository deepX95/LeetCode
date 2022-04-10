package simple

import "strings"

// rotateString 796. 旋转字符串
// fixme:字符串匹配算法（kmp）
func rotateString(s string, goal string) bool {
	// 自己写一个kmp算法

	// 调用库函数（kmp）
	str := strings.Builder{}
	str.WriteString(goal)
	str.WriteString(goal)

	return strings.Contains(str.String(), s)
}

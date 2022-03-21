package simple

import "strings"

// replaceSpace 剑指 Offer 05. 替换空格
func replaceSpace(s string) string {
	// 遍历
	res := strings.Builder{}
	for i := range s {
		switch s[i] {
		case ' ':
			res.WriteString("%20")
		default:
			res.WriteByte(s[i])
		}
	}
	return res.String()

	// 库函数
	//return strings.Replace(s," ","%20",-1)
}

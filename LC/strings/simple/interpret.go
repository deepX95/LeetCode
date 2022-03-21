package simple

import "strings"

// interpret 1678. 设计 Goal 解析器
func interpret(command string) string {
	res := strings.Builder{}
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res.WriteByte(command[i])
		} else if command[i] == ')' && command[i-1] == '(' {
			res.WriteByte('o')
		} else if command[i] == ')' && command[i-1] == 'l' {
			res.WriteString("al")
		}
	}
	return res.String()
}

package simple

// removeDuplicates 1047. 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

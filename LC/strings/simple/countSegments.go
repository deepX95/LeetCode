package simple

// countSegments 434. 字符串中的单词数
func countSegments(s string) int {
	cnt := 0
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			cnt++
		}
	}
	return cnt
}

package simple

// truncateSentence 1816. 截断句子
func truncateSentence(s string, k int) string {
	cnt := 0
	for i := range s {
		if i+1 == len(s) || s[i+1] == ' ' {
			cnt++
		}
		if cnt == k {
			return s[:i+1]
		}
	}
	return ""
}

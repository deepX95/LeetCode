package simple

// checkIfPangram 1832. 判断句子是否为全字母句
// fixme:位运算
func checkIfPangram(sentence string) bool {
	m := 0
	for _, b := range sentence {
		m |= 1 << (b - 'a')
	}
	return m == 1<<26-1
}

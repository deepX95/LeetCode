package simple

// isUnique 面试题 01.01. 判定字符是否唯一
func isUnique(astr string) bool {
	var mark = 0
	for _, c := range astr{
		bit := c - 'a'
		if (mark & (1 << bit)) != 0{
			return false
		}
		mark |= 1 << bit
	}
	return true
}

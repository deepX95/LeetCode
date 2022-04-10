package simple

// cellsInRange 2194. Excel 表中某个范围内的单元格
func cellsInRange(s string) []string {
	ret := make([]string, 0, (s[3]-s[0]+1)*(s[4]-s[1]+1))
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			ret = append(ret, string(i)+string(j))
		}
	}
	return ret
}

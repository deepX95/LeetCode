package simple

// firstUniqChar 387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	cnt:=[26]int{}
	for i:=range s{
		cnt[s[i]-'a']++
	}
	for i:=range s{
		if cnt[s[i]-'a']==1{
			return i
		}
	}
	return -1
}

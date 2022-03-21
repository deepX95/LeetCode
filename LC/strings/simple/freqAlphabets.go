package simple

import "strings"

// freqAlphabets 1309. 解码字母到整数映射
func freqAlphabets(s string) string {
	res := strings.Builder{}
	n := len(s)
	for i := 0; i < n; i++ {
		if i+2 < n && s[i+2] == '#' {
			res.WriteByte((s[i]-'0')*10 + (s[i+1] - '1') + 'a')
			i += 2
		} else {
			res.WriteByte(s[i] - '1' + 'a')
		}
	}
	return res.String()
}

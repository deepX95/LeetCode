package simple

import (
	"bytes"
)

// removeOuterParentheses 1021. 删除最外层的括号
func removeOuterParentheses(s string) string {
	/*
	( () () ) ()

	 */
	ret,cnt:=bytes.Buffer{},0
	for i:=0;i<len(s);i++{
		if s[i]=='(' {
			if cnt>0 {
				ret.WriteByte(s[i])
			}
			cnt++
		}else if s[i]==')' && cnt>1{
			if cnt>1 {
				ret.WriteByte(s[i])
			}
			cnt--
		}
	}
	return ret.String()
}

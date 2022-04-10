package simple

// maxDepth 1614. 括号的最大嵌套深度
func maxDepth(s string) int {
	max,cnt:=0,0
	for i:=range s{
		if s[i]=='('{
			cnt++
			if cnt>max{
				max=cnt
			}
		}else if s[i]==')'{
			cnt--
		}
	}
	return max
}

package simple

// balancedStringSplit
func balancedStringSplit(s string) int {
	cnt,flag:=0,0
	for i:=range s{
		if s[i]=='R'{
			flag++
		}else{
			flag--
		}

		if flag==0{
			cnt++
		}
	}
	return cnt
}

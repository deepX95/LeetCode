package simple

// romanToInt 13. 罗马数字转整数
func romanToInt(s string)(ret int) {
	symbolValues := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,
		'D': 500, 'M': 1000}

	// 法 1
	//ret=symbolValues[s[0]]
	//for i:=1;i<len(s);i++{
	//	pre := symbolValues[s[i - 1]]
	//	cur := symbolValues[s[i]]
	//	if cur>pre{
	//		ret+=cur-2*pre
	//	}else{
	//		ret+=cur
	//	}
	//}
	//return

	// 法2
	n:=len(s)
	for i:=0;i<n-1;i++{
		if symbolValues[s[i]]<symbolValues[s[i+1]]{
			ret-= symbolValues[s[i]]
		}else{
			ret+= symbolValues[s[i]]
		}
	}
	ret+= symbolValues[s[n-1]]
	return
}

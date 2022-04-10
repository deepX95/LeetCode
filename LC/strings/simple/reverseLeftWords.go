package simple

// reverseLeftWords 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	byt:= stringToBytes(s)
	cnt:=0
	for i:=0;i<len(s);i++{
		cnt++
		if cnt==n{
			tmp:=make([]byte,n)
			copy(tmp,byt[:n])
			copy(byt[:len(s)-n],byt[n:])
			copy(byt[len(s)-n:],tmp)
			break
		}
	}

	return string(byt)
}

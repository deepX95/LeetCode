package simple

// reverseStr 541. 反转字符串 II
func reverseStr(s string, k int) string {
	t := []byte(s)
	cnt:=0
	for i:=range s{
		cnt++
		if cnt==k{
			tmp:=reverse(t[i+1-k:i+1])
			copy(t[i+1-k:i+1],tmp)
			cnt=0
		}
	}

	return string(t)

	//t := []byte(s)
	//for i := 0; i < len(s); i += 2 * k {
	//	sub := t[i:min(i+k, len(s))]
	//	for j, n := 0, len(sub); j < n/2; j++ {
	//		sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
	//	}
	//}
	//return string(t)
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func reverse(t []byte)[]byte{
	//defer func() {
	//	if err:=recover();err!=nil{
	//		s:="s"
	//		fmt.Println(err,s)
	//	}
	//}()

	//l,r:=0,k-1
	//for l<r{
	//	bytes[l],bytes[r]=bytes[r],bytes[l]
	//	l++
	//	r--
	//}
	//return bytes
	for left, right := 0, len(t)-1; left < right; left++ {
		t[left], t[right] = t[right], t[left]
		right--
	}
	return t
}


//func stringToBytes(str string) []byte {
//	return *(*[]byte)(unsafe.Pointer(
//		&struct {
//			string
//			Cap int
//		}{str, len(str)},
//	))
//}

package simple

import "strconv"

// addBinary 67. 二进制求和
// fixme:再写一遍
func addBinary(a string, b string) (ret string) {
	carry:=0
	la,lb:=len(a),len(b)
	n:=max1(la,lb)

	for i:=0;i<n;i++{
		if i<la{
			carry+=int(a[la-i-1]-'0')
		}
		if i<lb{
			carry+=int(b[lb-i-1]-'0')
		}
		ret=strconv.Itoa(carry%2)+ret
		carry/=2
	}
	if carry>0{
		ret="1"+ret
	}

	return
}

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}
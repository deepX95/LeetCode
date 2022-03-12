package simple

// decode 1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	// 0010
	// 0001
	// 0011

	// 0100
	// 0110
	// 0010

	ret:=make([]int,len(encoded)+1)
	ret[0]=first
	for i:=range encoded{
		ret[i+1]=ret[i]^encoded[i]
	}
	return ret
}

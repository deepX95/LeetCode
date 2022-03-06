package simple

// maximum 面试题 16.07. 最大数值
// fixme:不会，再看再看
func maximum(a int, b int) int {
	// 0000 0001
	// 0000 0010

	//整数右移高位补0,负数右移高位补1
	//a>b,ret>0,temp&1=0
	//a<b,ret<0,temp&1=1
	ret := int64(a-b)
	tmp:=ret&(ret>>63)
	ret = int64(a)-tmp
	return int(ret)
}

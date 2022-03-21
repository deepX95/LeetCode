package simple

// exchangeBits 面试题 05.07. 配对交换
// fixme:好好品味下
func exchangeBits(num int) int {
	// 0101 0101
	// 1010 1010

	// 0000 0010
	// 0000 0001

	// 0000 0000
	// 0000 0001

	// 奇数位向左移一位 偶数位向右移一位  对以上结果或运算
	//return ((num & 0x55555555) << 1) | ((num & 0xaaaaaaaa) >> 1)

	// 模拟
	i,j:=0,1
	for i<=30{
		a:=num>>i&1
		b:=num>>j&1
		if a!=b{
			num^=1<<i
			num^=1<<j
		}
		i+=2
		j+=2
	}
	return num
}

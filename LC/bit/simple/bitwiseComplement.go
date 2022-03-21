package simple

// bitwiseComplement 1009. 十进制整数的反码
func bitwiseComplement(n int) int {
	// 0000 0101   5
	// 0000 0010   2
	// 0000 0111   7

	// 0000 0001   1
	// 0000 0011   3
	// 0000 0111   7

	// fixme:https://leetcode-cn.com/problems/complement-of-base-10-integer/solution/1chong-jie-fa-2ge-jiao-du-2fen-ji-jian-d-wlpo/
	// 前导置1
	if n==0{
		return 1
	}
	mask:= ^0
	for mask&n>0{
		mask<<=1
	}
	return mask ^^n


	// 后导置1
	//bit:=1
	//for bit<n{
	//	bit=(bit<<1)+1
	//}
	//return bit^n

}

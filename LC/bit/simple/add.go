package simple

// add 面试题 17.01. 不用加号的加法
// fixme:好好再看一遍
// fixme:https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/solution/dian-zan-yao-wo-zhi-dao-ni-xiang-kan-dia-ovxy/
func add(a int, b int) int {
	// 0000 1001  9
	// 0000 0010  2
	// 0000 1011  11

	//    1101111
	// 1110000011
	//
	// 1111110010

	for b!=0{  // 当进位为 0 时跳出
		carry:=a&b // 计算 进位   计算出2个加数二进制下每一位的进位
		a ^= b  // 计算 本位   计算出2个加数二进制下每一位的本位
		b = carry<<1 // 进位做进位逻辑，也就是 * 2
	}
	return a
}

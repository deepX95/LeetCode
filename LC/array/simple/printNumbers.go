package simple

import "math"

// printNumbers 剑指 Offer 17. 打印从1到最大的n位数
// fixme:大数问题，之前回溯熟练了再看此题
func printNumbers(n int) []int {
	// 法1
	lp := int(math.Pow10(n))
	ret := make([]int, 0, lp)
	for i := 1; i < lp; i++ {
		ret = append(ret, i)
	}
	return ret


}

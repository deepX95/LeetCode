package simple

// findTheDifference 389. 找不同
func findTheDifference(s string, t string) (diff byte) {
	// 位运算
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]

	// 求和
	//sum := 0
	//for _, ch := range s {
	//	sum -= int(ch)
	//}
	//for _, ch := range t {
	//	sum += int(ch)
	//}
	//return byte(sum)

	// 计算
	//cnt := [26]int{}
	//for _, ch := range s {
	//	cnt[ch-'a']++
	//}
	//for i := 0; ; i++ {
	//	ch := t[i]
	//	cnt[ch-'a']--
	//	if cnt[ch-'a'] < 0 {
	//		return ch
	//	}
	//}
}

package simple

// lengthOfLongestSubstring 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m:=map[byte]int{}
	n:=len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk,ans:=-1,0
	for i:=0;i<n;i++{
		if i!=0{
			// 左指针向右移动一格，移除一个字符
			delete(m,s[i-1])
		}
		for rk+1<n && m[s[rk+1]]==0{
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = maxLengthOfLongestSubstring(ans, rk - i + 1)
	}
	return ans

	// fixme:https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/longest-substring-without-repeating-characters-b-2/
	//m := make([]int, 128)
	//ans := 0
	//i := 0
	//for j := 0; j < len(s); j++ {
	//	i = maxLengthOfLongestSubstring(i, m[s[j]])
	//	m[s[j]] = j + 1
	//	ans = maxLengthOfLongestSubstring(ans, j-i+1)
	//}
	//return ans
}

func maxLengthOfLongestSubstring(a, b int) int {
	if a > b {
		return a
	}
	return b
}

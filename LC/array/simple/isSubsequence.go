package simple

// isSubsequence 392. 判断子序列
// fixme:动态规划 https://leetcode-cn.com/problems/is-subsequence/solution/pan-duan-zi-xu-lie-by-leetcode-solution/
func isSubsequence(s string, t string) bool {
	// 双指针
	//if(s == "" && t!="")  || (s=="" && t=="") {
	//	return true
	//}
   //
	//if s!="" && t==""{
	//	return false
	//}
   //
   //l,r:=0,len(s)
   //rr:=0
   //for l<r && rr<len(t){
   //	 if s[l]==t[rr]{
   //	 	l++
   //	 	if l==r{
   //	 		return true
	//	}
	// }
	// rr++
   //}
   //return false

   // 双指针
	n,m:=len(s),len(t)
	i,j:=0,0
	for i<n && j<m{
		if s[i]==t[j]{
			i++
		}
		j++
	}
	return i==n

	// 动态规划
	//n, m := len(s), len(t)
	//f := make([][26]int, m + 1)
	//for i := 0; i < 26; i++ {
	//	f[m][i] = m
	//}
	//for i := m - 1; i >= 0; i-- {
	//	for j := 0; j < 26; j++ {
	//		if t[i] == byte(j + 'a') {
	//			f[i][j] = i
	//		} else {
	//			f[i][j] = f[i + 1][j]
	//		}
	//	}
	//}
	//add := 0
	//for i := 0; i < n; i++ {
	//	if f[add][int(s[i] - 'a')] == m {
	//		return false
	//	}
	//	add = f[add][int(s[i] - 'a')] + 1
	//}
	//return true

}
package simple

// countConsistentStrings 1684. 统计一致字符串的数目
func countConsistentStrings(allowed string, words []string) int {
	//	mp := make(map[byte]struct{}, len(allowed))
	//	for i := range allowed {
	//		mp[allowed[i]] = struct{}{}
	//	}
	//
	//	cnt := 0
	//ct:
	//	for i := range words {
	//		for j := range words[i] {
	//			if _, ok := mp[words[i][j]]; !ok {
	//				continue ct
	//			}
	//		}
	//		cnt++
	//	}
	//	return cnt

	var allowedStringMap, count = [26]bool{}, 0
	for _, char := range allowed {
		allowedStringMap[char-'a'] = true
	}

	for _, word := range words {
		isUnanimousWord := true
		for _, char := range word {
			if !allowedStringMap[char-'a'] {
				isUnanimousWord = false
				break
			}
		}
		if isUnanimousWord {
			count++
		}
	}
	return count
}

package simple

// buildArray_2 1441. 用栈操作构建数组
func buildArray_2(target []int, n int) []string {
	//l := len(target)
	//mp := make(map[int]struct{}, len(target))
	//for i := range target {
	//	mp[target[i]] = struct{}{}
	//}
	//
	//var ret []string
	//for i := 1; i <= n; i++ {
	//	if _, ok := mp[i]; ok {
	//		ret=append(ret,"Push")
	//		if i == target[l-1] {
	//			break
	//		}
	//		continue
	//	}
	//	ret=append(ret,"Push","Pop")
	//}
	//return ret

	var size int
	var prev int
	for _, val := range target {
		size += (val - prev - 1) * 2 + 1
		prev = val
	}

	res := make([]string, 0, size)
	prev = 0
	for _, val := range target {
		dropCount := val - prev - 1
		for i := 0; i < dropCount; i++ {
			res = append(res, "Push", "Pop")
		}
		res = append(res, "Push")
		prev = val
	}

	return res
}

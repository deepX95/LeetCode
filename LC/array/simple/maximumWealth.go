package simple

// maximumWealth 1672. 最富有客户的资产总量
func maximumWealth(accounts [][]int) int {
	ret := 0
	for i := range accounts {
		tmp := sumInt(accounts[i])
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

func sumInt(in []int) (ret int) {
	for i := range in {
		ret += in[i]
	}
	return
}

package simple

// numJewelsInStones 771. 宝石与石头
func numJewelsInStones(jewels string, stones string) int {
	mp := make(map[byte]struct{}, len(jewels))
	for i := range jewels {
		mp[jewels[i]] = struct{}{}
	}

	cnt := 0
	for i := range stones {
		if _, ok := mp[stones[i]]; ok {
			cnt++
		}
	}

	return cnt
}

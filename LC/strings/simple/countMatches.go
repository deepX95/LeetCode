package simple

// countMatches 1773. 统计匹配检索规则的物品数量
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	cnt:=0
	for i:=range items{
		switch ruleKey {
		case "type":
			if items[i][0]==ruleValue{
				cnt++
			}
		case "color":
			if items[i][1]==ruleValue{
				cnt++
			}
		case "name":
			if items[i][2]==ruleValue{
				cnt++
			}
		}
	}
	return cnt
}
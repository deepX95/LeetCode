package simple

// destCity 1436. 旅行终点站
func destCity(paths [][]string) string {
	mp := make(map[string]string, len(paths))
	for i := range paths {
		mp[paths[i][0]] = paths[i][1]
	}
	for i:=range paths{
		dst:=mp[paths[i][0]]
		if _,ok:=mp[dst];!ok{
			return dst
		}
	}
	return ""
}

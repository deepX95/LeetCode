package simple

import "strings"

// defangIPaddr  1108. IP 地址无效化
func defangIPaddr(address string) string {
	//return strings.Replace(address,".","[.]",-1)
	ret:=strings.Builder{}
	for  i :=range address{
		if address[i]=='.' {
			ret.WriteString("[.]")
		}else {
			ret.WriteByte(address[i])
		}
	}
	return ret.String()
}

package simple

// numberOfSteps 1342. 将数字变成 0 的操作次数
func numberOfSteps(num int) (cnt int) {
     // 0000 1000
	for num!=0{
		if num&1==0{
			num>>=1
		}else{
			num-=1
		}
		cnt++
	}
	return cnt

	//for ; num > 0; num >>= 1 {
	//	cnt += num & 1
	//	if num > 1 {
	//		cnt++
	//	}
	//}
	//return

}

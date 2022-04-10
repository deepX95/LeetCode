package simple

// numberOfMatchess 1688. 比赛中的配对次数
func numberOfMatches(n int) (ret int) {
	for n>1{
		if n%2==0{
			ret+=n/2
			n/=2
		}else{
			ret+=(n-1)/2
			n=(n-1)/2+1
		}
	}
	return 0
}

package simple


// MySqrt 69. x 的平方根
func MySqrt(x int) int {
	l,r:=0,x
	var ret int
	for l<=r{
		mid:=l+(r-l)/2
		if mid*mid<=x{
			ret=mid
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return ret
}
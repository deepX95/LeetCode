package simple

// IsPerfectSquare  367. 有效的完全平方数
func IsPerfectSquare(num int) bool {
	l,r:=0,num
	for l<=r{
		mid:=l+(r-l)/2
		if mid*mid==num{
			return true
		}else if mid*mid>num{
			r=mid-1
		}else{
			l=mid+1
		}
	}
	return false
}
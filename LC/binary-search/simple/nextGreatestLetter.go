package simple

// nextGreatestLetter 744. 寻找比目标字母大的最小字母
func nextGreatestLetter(letters []byte, target byte) byte {
	l,r,mid:=0,len(letters)-1,0

	for l<r{
		mid=l+(r-l)/2
		if letters[mid] <=target{
			l=mid+1
		}else{
			r=mid
		}
	}
	if l==len(letters){
		return letters[0]
	}

	return letters[l]
}

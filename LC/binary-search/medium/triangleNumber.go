package medium

import "sort"

// triangleNumber 611. 有效三角形的个数
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	lth,count,mid:=len(nums),0,0
	for i:=0;i<lth-2;i++{
		for j:=i+1;j<lth-1;j++{
			// 找到第1个大于等于两边之和的下标
			l,r:=j+1,lth
			for l<r{
				mid=l+(r-l)/2
				if nums[mid]<nums[i]+nums[j]{
					l=mid+1
				}else{
					r=mid
				}
			}
			count+=l-j-1
		}
	}
	return count
}
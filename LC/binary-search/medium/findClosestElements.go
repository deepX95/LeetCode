package medium

// findClosestElements 658. 找到 K 个最接近的元素
// fixme:重点再做几次
func findClosestElements(arr []int, k int, x int) []int {
	l, r, mid := 0, len(arr)-k, 0
	for l < r {
		mid = l + (r-l)/2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}

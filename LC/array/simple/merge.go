package simple

// merge 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	l, m, n := m+n-1, m-1, n-1
	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			m--
		} else {
			nums1[l] = nums2[n]
			n--
		}
		l--
	}
}

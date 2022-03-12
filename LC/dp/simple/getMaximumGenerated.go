package simple

// getMaximumGenerated 1646. 获取生成数组中的最大值

func getMaximumGenerated(n int) (ans int) {
	nums := make([]int, n+1)
	nums[1] = 0
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
		//if i%2 == 0 {
		//	nums[i] = nums[i/2]
		//} else {
		//	nums[i] = nums[i/2] + nums[i/2+1]
		//}

	}
	for i := range nums {
		ans = maxMaximumGenerated(ans, nums[i])
	}
	return
}

func maxMaximumGenerated(a, b int) int {
	if a > b {
		return a
	}
	return b
}

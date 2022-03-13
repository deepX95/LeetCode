package simple

// kidsWithCandies 1431. 拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for i := 0; i < n; i++ {
		maxCandies = maxKidswithcandies(maxCandies, candies[i])
	}
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = candies[i]+extraCandies >= maxCandies
	}
	return ret
}

func maxKidswithcandies(x, y int) int {
	if x > y {
		return x
	}
	return y
}

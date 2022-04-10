package simple

// countPoints 2103. 环和杆
func countPoints(rings string) int {
	n := len(rings)
	status := make([]int, 10)
	for i := 0; i < n; i += 2 {
		idx := rings[i+1] - '0'
		switch rings[i] {
		case 'R':
			status[idx] |= 1
		case 'G':
			status[idx] |= 2
		default:
			status[idx] |= 4
		}
	}
	ret:=0
	for i:=0;i<10;i++{
		if status[i]==7{
			ret++
		}
	}
	return ret
}

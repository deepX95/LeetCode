package simple

// finalValueAfterOperations 2011.执行操作后的变量值
func finalValueAfterOperations(operations []string) int {
	x := 0
	for i := range operations {
		if operations[i] == "--X" || operations[i] == "X--" {
			x--
			continue
		}
		x++
	}
	return x
}

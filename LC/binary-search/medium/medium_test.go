package medium

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5,7,7,8,8,10}
	ret := searchRange(nums, 8)
	fmt.Println(ret)
}

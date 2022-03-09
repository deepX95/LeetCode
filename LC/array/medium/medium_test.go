package medium

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T){
	singleNumber([]int{2,2,3,2})
}

func TestPermute(t *testing.T){
	ret:=permute([]int{1,2})
	fmt.Println(ret)
}
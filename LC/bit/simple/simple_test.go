package simple

import (
	"fmt"
	"testing"
)

func TestIsUnique(t *testing.T) {
	isUnique("aabc")
}

func TestMissingNumber(t *testing.T) {
	ret := missingNumber([]int{0})
	fmt.Println(ret)
}

func TestCountBits(t *testing.T) {
	countBits(2)
}

func TestBinaryGap(t *testing.T) {
	binaryGap(8)
}

func TestConvertInteger(t *testing.T) {
	convertInteger(29, 15)
}

func TestAdd(t *testing.T) {
	add(9, 2)
}

func TestNumberOfSteps(t *testing.T) {
	numberOfSteps(123)
}

func TestHasAlternatingBits(t *testing.T) {
	hasAlternatingBits(10)
}

func TestBitwiseComplement(t *testing.T){
	bitwiseComplement(5)
}

func TestDecode(t *testing.T)  {
	decode([]int{1,2,3},1)
}
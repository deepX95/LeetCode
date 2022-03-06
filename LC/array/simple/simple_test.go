package simple

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				nums: []int{1, 2, 0, 0, 3, 4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.args.nums)
		})
	}
}

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				nums: []int{2, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	ret := romanToInt("MCMXCIV")
	fmt.Println(ret)
}

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{1, 1}
	ret := findDisappearedNumbers(nums)
	fmt.Println(ret)
}
func TestAddBinary(t *testing.T) {
	addBinary("11", "1")
}

func TestMerge(t *testing.T) {
	//merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0}, 0, []int{1}, 1)
	//merge([]int{4,5,6,0,0,0}, 3, []int{1,2,3}, 3)

}

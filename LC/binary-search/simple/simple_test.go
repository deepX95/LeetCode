package simple

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				nums: []int{
					-1, 0, 3, 5, 9,
				},
				target: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	ret := searchInsert(nums, 7)
	fmt.Println(ret)
}


func TestNextGreatestLetter(t *testing.T) {
	letters  := []byte{'c', 'f', 'j'}
	ret := nextGreatestLetter(letters , 'j')
	fmt.Println(ret)
}


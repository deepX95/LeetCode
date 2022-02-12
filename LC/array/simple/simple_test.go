package simple

import "testing"

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

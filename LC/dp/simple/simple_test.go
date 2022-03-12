package simple

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				1,
			},
		},
		{
			name: "2",
			args: args{
				2,
			},
		},
		{
			name: "3",
			args: args{
				3,
			},
		},
		{
			name: "4",
			args: args{
				4,
			},
		},
		{
			name: "5",
			args: args{
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret := maxSubArray(nums)
	fmt.Println(ret)
}
func Test_tribonacci(t *testing.T) {
	ret:=tribonacci(25)
	println(ret)
}

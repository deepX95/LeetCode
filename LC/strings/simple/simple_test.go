package simple

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				s: []byte{
					'h', 'e', 'l', 'l', 'o',
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.args.s)
		})
	}
}

func Test_toLowerCase(t *testing.T) {
	toLowerCase("Hello")
}

func Test_isFlipedString(t *testing.T) {
	isFlipedString("aba", "bab")
}

func Test_freqAlphabets(t *testing.T) {
	ch := '1' - '1' + 'a'
	fmt.Println(string(ch))
}

func Test_lengthOfLongestSubstring(t *testing.T){
	lengthOfLongestSubstring("aabja")
}

func Test_calculate(t *testing.T){
	calculate("AB")
}
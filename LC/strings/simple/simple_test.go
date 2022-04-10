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

func Test_reverseStr(t *testing.T){
	reverseStr("abcdefg",2)
}

func Test_reversePrefix(t *testing.T){
	reversePrefix("abcdefd",'d')
}

func Test_lengthOfLongestSubstring(t *testing.T){
	lengthOfLongestSubstring("aabja")
}

func Test_calculate(t *testing.T){
	calculate("AB")
}

func Test_maxDepth(t *testing.T){
	maxDepth("(1+(2*3)+((8)/4))+1")
}

func Test_reverseLeftWords(t *testing.T){
	reverseLeftWords("abcdefg",2)
}
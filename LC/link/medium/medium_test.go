package medium

import (
	"LeetCode/LC/link"
	"reflect"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	type args struct {
		list1 *link.ListNode
		a     int
		b     int
		list2 *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name:"test",
			args: args{
				list1: &link.ListNode{
					Val: 0,
					Next: &link.ListNode{
						Val: 1,
						Next:  &link.ListNode{
							Val:2,
								Next:  &link.ListNode{
									Val: 3,
										Next:  &link.ListNode{
											Val: 4,
												Next:  &link.ListNode{
													Val: 5,
														Next: nil},
										}}}}},
				a:3,
				b:4,
				list2: &link.ListNode{
				Val: 	1000000,
				Next: &link.ListNode{
				Val: 1000001,
				Next: &link.ListNode{
					Val: 1000002,
					Next: nil,
				},
				},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

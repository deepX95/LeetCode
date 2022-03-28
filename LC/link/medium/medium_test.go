package medium

import (
	"LeetCode/LC/link"
	"fmt"
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
			name: "test",
			args: args{
				list1: &link.ListNode{
					Val: 0,
					Next: &link.ListNode{
						Val: 1,
						Next: &link.ListNode{
							Val: 2,
							Next: &link.ListNode{
								Val: 3,
								Next: &link.ListNode{
									Val: 4,
									Next: &link.ListNode{
										Val:  5,
										Next: nil},
								}}}}},
				a: 3,
				b: 4,
				list2: &link.ListNode{
					Val: 1000000,
					Next: &link.ListNode{
						Val: 1000001,
						Next: &link.ListNode{
							Val:  1000002,
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

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *link.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "many",
			args: args{
				head: &link.ListNode{
					Val: 1,
					Next: &link.ListNode{
						Val: 2,
						Next: &link.ListNode{
							Val: 3,
							Next: &link.ListNode{
								Val: 4,
								Next: &link.ListNode{
									Val:  5,
									Next: nil},
							}}}},
				n: 5,
			},
		},
		{
			name: "one",
			args: args{
				head: &link.ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test",
			args: args{
				head: &link.ListNode{
					Val: 1,
					Next: &link.ListNode{
						Val: 1,
						Next: &link.ListNode{
							Val: 1,
							Next: &link.ListNode{
								Val: 2,
								Next: &link.ListNode{
									Val:  3,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList(t *testing.T) {
	/**
	 * Your MyLinkedList object will be instantiated and called as such:
	 * obj := Constructor();
	 * param_1 := obj.Get(index);
	 * obj.AddAtHead(val);
	 * obj.AddAtTail(val);
	 * obj.AddAtIndex(index,val);
	 * obj.DeleteAtIndex(index);
	 */
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(0)
	fmt.Println(obj.Get(0))

	//obj := Constructor()
	//obj.AddAtIndex(0, 10)
	//obj.AddAtIndex(0, 20)
	//obj.AddAtIndex(1, 30)
	//fmt.Println(obj.Get(0))

}

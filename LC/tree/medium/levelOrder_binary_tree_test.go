package medium

import (
	"LeetCode/LC/tree"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name:"test",
			args: args{
				root: &tree.TreeNode{
					Val: 3,
					Left: &tree.TreeNode{
						Val: 9,
					},
					Right: &tree.TreeNode{
						Val: 20,
						Left: &tree.TreeNode{
							Val: 15,
						},
						Right: &tree.TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

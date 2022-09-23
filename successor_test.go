package avltree

import (
	"reflect"
	"testing"
)

func TestAVLTree_inorderSuccessorByNode(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	type args struct {
		subTreeKey int
		nodeKey    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantKey int
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 12,
					Left: &Node[int, string]{
						Key: 7,
						Left: &Node[int, string]{
							Key: 6,
						},
						Right: &Node[int, string]{
							Key: 10,
						},
					},
					Right: &Node[int, string]{
						Key: 15,
						Left: &Node[int, string]{
							Key: 13,
							Right: &Node[int, string]{
								Key: 14,
							},
						},
						Right: &Node[int, string]{
							Key: 20,
						},
					},
				},
				size:      0,
				threshold: 1,
			},
			args: args{
				subTreeKey: 12,
				nodeKey:    14,
			},
			wantKey: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}

			subTree := tree.search(tt.args.subTreeKey)
			want := tree.search(tt.wantKey)

			if got := tree.inorderSuccessor(subTree, tt.args.nodeKey); !reflect.DeepEqual(got, want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, want)
			}
		})
	}
}

func TestAVLTree_InorderSuccessor(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantK  int
		wantV  string
		wantOK bool
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 12,
					Left: &Node[int, string]{
						Key: 7,
						Left: &Node[int, string]{
							Key: 6,
						},
						Right: &Node[int, string]{
							Key: 10,
						},
					},
					Right: &Node[int, string]{
						Key:   15,
						Value: "correct",
						Left: &Node[int, string]{
							Key: 13,
							Right: &Node[int, string]{
								Key: 14,
							},
						},
						Right: &Node[int, string]{
							Key: 20,
						},
					},
				},
			},
			args: args{
				key: 14,
			},
			wantK:  15,
			wantV:  "correct",
			wantOK: true,
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 12,
					Left: &Node[int, string]{
						Key: 7,
						Left: &Node[int, string]{
							Key: 6,
						},
						Right: &Node[int, string]{
							Key:   10,
							Value: "correct",
						},
					},
					Right: &Node[int, string]{
						Key: 15,
						Left: &Node[int, string]{
							Key: 13,
							Right: &Node[int, string]{
								Key: 14,
							},
						},
						Right: &Node[int, string]{
							Key: 20,
						},
					},
				},
				size:      8,
				threshold: 1,
			},
			args: args{
				key: 7,
			},
			wantK:  10,
			wantV:  "correct",
			wantOK: true,
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 12,
					Left: &Node[int, string]{
						Key: 7,
						Left: &Node[int, string]{
							Key: 6,
						},
						Right: &Node[int, string]{
							Key: 10,
						},
						Height: 1,
					},
					Right: &Node[int, string]{
						Key:   15,
						Value: "correct",
						Left: &Node[int, string]{
							Key: 13,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 20,
						},
						Height: 2,
					},
				},
				size:      8,
				threshold: 1,
			},
			args: args{
				key: 20,
			},
			wantK:  0,
			wantV:  "",
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}
			gotK, gotV, gotOK := tree.InorderSuccessor(tt.args.key)
			if gotOK != tt.wantOK {
				t.Fatalf("InorderSuccessor() gotOK = %v, wantOK %v", gotOK, tt.wantOK)
			}
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Fatalf("InorderSuccessor() gotK = %v, wantK %v", gotK, tt.wantK)
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("InorderSuccessor() gotV = %v, wantV %v", gotV, tt.wantV)
			}
		})
	}
}

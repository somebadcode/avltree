package avltree

import (
	"reflect"
	"testing"
)

func TestAVLTree_Delete(t *testing.T) {
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
		want   *Node[int, string]
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
						Height: 1,
					},
					Right: &Node[int, string]{
						Key: 15,
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
					Height: 3,
				},
				size:      8,
				threshold: 1,
			},
			args: args{
				key: 7,
			},
			want: &Node[int, string]{
				Key: 12,
				Left: &Node[int, string]{
					Key: 10,
					Left: &Node[int, string]{
						Key: 6,
					},
					Height: 1,
				},
				Right: &Node[int, string]{
					Key: 15,
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
				Height: 3,
			},
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
						Key: 15,
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
					Height: 3,
				},
				size:      8,
				threshold: 1,
			},
			args: args{
				key: 13,
			},
			want: &Node[int, string]{
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
					Key: 15,
					Left: &Node[int, string]{
						Key: 14,
					},
					Right: &Node[int, string]{
						Key: 20,
					},
					Height: 2,
				},
				Height: 3,
			},
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 12,
					Left: &Node[int, string]{
						Key: 7,
						Left: &Node[int, string]{
							Key: 6,
							Left: &Node[int, string]{
								Key: 5,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 10,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 15,
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
					Height: 3,
				},
				size:      8,
				threshold: 1,
			},
			args: args{
				key: 6,
			},
			want: &Node[int, string]{
				Key: 12,
				Left: &Node[int, string]{
					Key: 7,
					Left: &Node[int, string]{
						Key: 5,
					},
					Right: &Node[int, string]{
						Key: 10,
					},
					Height: 2,
				},
				Right: &Node[int, string]{
					Key: 15,
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
				Height: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}

			tree.Delete(tt.args.key)

			if !reflect.DeepEqual(tree.root, tt.want) {
				t.Errorf("Delete() = %v, want %v", tree.root, tt.want)
			}
		})
	}
}

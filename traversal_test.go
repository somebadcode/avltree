package avltree

import (
	"reflect"
	"testing"
)

func TestAVLTree_InorderTraversal(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	tests := []struct {
		name      string
		fields    fields
		stop      bool
		stopAtKey int
		want      []int
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			want: []int{2, 14, 20, 23, 33, 90, 99, 300, 777, 900, 920},
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			stop:      true,
			stopAtKey: 90,
			want:      []int{2, 14, 20, 23, 33, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}

			got := make([]int, 0, tree.size)

			tree.InorderTraversal(func(key int, _ string) bool {
				got = append(got, key)

				return !(tt.stop && key == tt.stopAtKey)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_PostorderTraversal(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	tests := []struct {
		name      string
		fields    fields
		stop      bool
		stopAtKey int
		want      []int
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			want: []int{14, 2, 23, 20, 99, 90, 777, 920, 900, 300, 33},
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			stop:      true,
			stopAtKey: 90,
			want:      []int{14, 2, 23, 20, 99, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}

			got := make([]int, 0, tree.size)

			tree.PostorderTraversal(func(key int, _ string) bool {
				got = append(got, key)

				return !(tt.stop && key == tt.stopAtKey)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_PreorderTraversal(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	tests := []struct {
		name      string
		fields    fields
		stop      bool
		stopAtKey int
		want      []int
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			want: []int{33, 20, 2, 14, 23, 300, 90, 99, 900, 777, 920},
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 33,
					Left: &Node[int, string]{
						Key: 20,
						Left: &Node[int, string]{
							Key: 2,
							Right: &Node[int, string]{
								Key: 14,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 23,
						},
						Height: 2,
					},
					Right: &Node[int, string]{
						Key: 300,
						Left: &Node[int, string]{
							Key: 90,
							Right: &Node[int, string]{
								Key: 99,
							},
							Height: 1,
						},
						Right: &Node[int, string]{
							Key: 900,
							Left: &Node[int, string]{
								Key: 777,
							},
							Right: &Node[int, string]{
								Key: 920,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				size:      11,
				threshold: 1,
			},
			stop:      true,
			stopAtKey: 23,
			want:      []int{33, 20, 2, 14, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}

			got := make([]int, 0, tree.size)

			tree.PreorderTraversal(func(key int, _ string) bool {
				got = append(got, key)

				return !(tt.stop && key == tt.stopAtKey)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

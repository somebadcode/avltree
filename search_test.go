package avltree

import (
	"reflect"
	"testing"
)

func TestAVLTree_Search(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantFound bool
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
			},
			args: args{
				key: 90,
			},
			wantFound: true,
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
			},
			args: args{
				key: 1,
			},
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}
			got, gotOK := tree.Search(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Search() got = %v, want %v", got, tt.want)
			}
			if gotOK != tt.wantFound {
				t.Errorf("Search() gotOK = %v, want %v", gotOK, tt.wantFound)
			}
		})
	}
}

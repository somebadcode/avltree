package avltree

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAVLTree_Insert(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	type args struct {
		key   int
		value string
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
					Key: 10,
					Left: &Node[int, string]{
						Key: 5,
					},
					Height: 1,
				},
				size:      2,
				threshold: 1,
			},
			args: args{
				key:   12,
				value: "",
			},
			want: &Node[int, string]{
				Key: 10,
				Left: &Node[int, string]{
					Key: 5,
				},
				Right: &Node[int, string]{
					Key: 12,
				},
				Height: 1,
			},
		},
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 10,
					Left: &Node[int, string]{
						Key: 5,
					},
					Height: 1,
				},
				size:      2,
				threshold: 1,
			},
			args: args{
				key:   5,
				value: "",
			},
			want: &Node[int, string]{
				Key: 10,
				Left: &Node[int, string]{
					Key: 5,
				},
				Height: 1,
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
			tree.Insert(tt.args.key, tt.args.value)

			if !reflect.DeepEqual(tree.root, tt.want) {
				rawJSON, err := json.Marshal(tt.want)
				if err != nil {
					t.Fatalf("Marshal() failed to marshal tt.want: %s", err)
				}

				wantJSON := string(rawJSON)

				rawJSON, err = json.Marshal(tree.root)
				if err != nil {
					t.Fatalf("Marshal() failed to marshal tree root: %s", err)
				}

				gotJSON := string(rawJSON)

				t.Errorf("Insert() = %v, want %v", gotJSON, wantJSON)
				t.Logf("Tree root balanace %d", balanceOf(tree.root))
			}
		})
	}
}

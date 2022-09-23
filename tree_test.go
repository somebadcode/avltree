package avltree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		threshold int
	}
	type insert struct {
		key   int
		value string
	}
	tests := []struct {
		name   string
		args   args
		insert []insert
		want   *AVLTree[int, string]
	}{
		{
			args: args{
				threshold: 1,
			},
			insert: []insert{
				{key: 1, value: "hello"},
				{key: 2, value: "world"},
				{key: 76, value: "more"},
				{key: 33, value: "hi"},
			},
			want: &AVLTree[int, string]{
				root: &Node[int, string]{
					Key:   2,
					Value: "world",
					Left: &Node[int, string]{
						Key:   1,
						Value: "hello",
					},
					Right: &Node[int, string]{
						Key:   76,
						Value: "more",
						Left: &Node[int, string]{
							Key:   33,
							Value: "hi",
						},
						Height: 1,
					},
					Height: 2,
				},
				size:      4,
				threshold: 1,
			},
		},
		{
			args: args{
				threshold: -5,
			},
			insert: []insert{
				{key: 1, value: "hello"},
				{key: 2, value: "world"},
				{key: 76, value: "more"},
				{key: 33, value: "hi"},
			},
			want: &AVLTree[int, string]{
				root: &Node[int, string]{
					Key:   2,
					Value: "world",
					Left: &Node[int, string]{
						Key:   1,
						Value: "hello",
					},
					Right: &Node[int, string]{
						Key:   76,
						Value: "more",
						Left: &Node[int, string]{
							Key:   33,
							Value: "hi",
						},
						Height: 1,
					},
					Height: 2,
				},
				size:      4,
				threshold: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New[int, string](tt.args.threshold)

			for _, kv := range tt.insert {
				tree.Insert(kv.key, kv.value)
			}

			if got := tree; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Size(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 5,
					Left: &Node[int, string]{
						Key: 2,
					},
					Height: 1,
				},
				size: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}
			if got := tree.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_RootKey(t *testing.T) {
	type fields struct {
		root      *Node[int, string]
		size      uint
		threshold int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				root: &Node[int, string]{
					Key: 50,
				},
				size:      1,
				threshold: 1,
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &AVLTree[int, string]{
				root:      tt.fields.root,
				size:      tt.fields.size,
				threshold: tt.fields.threshold,
			}
			if got := tree.RootKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

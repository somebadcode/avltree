package avltree

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_balance(t *testing.T) {
	type args struct {
		node      *Node[int, string]
		threshold int
	}
	tests := []struct {
		name string
		args args
		want *Node[int, string]
	}{
		{
			name: "left_heavy",
			args: args{
				node: &Node[int, string]{
					Key: 100,
					Left: &Node[int, string]{
						Key: 90,
						Left: &Node[int, string]{
							Key: 80,
						},
						Height: 1,
					},
					Height: 2,
				},
				threshold: 1,
			},
			want: &Node[int, string]{
				Key: 90,
				Left: &Node[int, string]{
					Key: 80,
				},
				Right: &Node[int, string]{
					Key: 100,
				},
				Height: 1,
			},
		},

		{
			name: "right_heavy",
			args: args{
				node: &Node[int, string]{
					Key: 100,
					Right: &Node[int, string]{
						Key: 110,
						Right: &Node[int, string]{
							Key: 120,
						},
						Height: 1,
					},
					Height: 2,
				},
				threshold: 1,
			},
			want: &Node[int, string]{
				Key: 110,
				Left: &Node[int, string]{
					Key: 100,
				},
				Right: &Node[int, string]{
					Key: 120,
				},
				Height: 1,
			},
		},
		{
			name: "left_right_heavy",
			args: args{
				node: &Node[int, string]{
					Key: 100,
					Left: &Node[int, string]{
						Key: 90,
						Right: &Node[int, string]{
							Key: 91,
							Right: &Node[int, string]{
								Key: 92,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				threshold: 1,
			},
			want: &Node[int, string]{
				Key: 91,
				Left: &Node[int, string]{
					Key: 90,
				},
				Right: &Node[int, string]{
					Key: 100,
					Left: &Node[int, string]{
						Key: 92,
					},
					Height: 1,
				},
				Height: 2,
			},
		},
		{
			name: "right_left_heavy",
			args: args{
				node: &Node[int, string]{
					Key: 100,
					Right: &Node[int, string]{
						Key: 120,
						Left: &Node[int, string]{
							Key: 125,
							Left: &Node[int, string]{
								Key: 124,
							},
							Height: 1,
						},
						Height: 2,
					},
					Height: 3,
				},
				threshold: 1,
			},
			want: &Node[int, string]{
				Key: 125,
				Left: &Node[int, string]{
					Key: 100,
					Right: &Node[int, string]{
						Key: 124,
					},
					Height: 1,
				},
				Right: &Node[int, string]{
					Key: 120,
				},
				Height: 2,
			},
		},
		{
			name: "nil",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balance(tt.args.node, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				rawJSON, err := json.Marshal(tt.want)
				if err != nil {
					t.Fatalf("Marshal() failed to marshal tt.want: %s", err)
				}

				wantJSON := string(rawJSON)

				rawJSON, err = json.Marshal(got)
				if err != nil {
					t.Fatalf("Marshal() failed to marshal tree root: %s", err)
				}

				gotJSON := string(rawJSON)

				t.Errorf("balance() = %v, want %v", gotJSON, wantJSON)
				t.Logf("Balanace: %d", balanceOf(got))
			}
		})
	}
}

package avltree

import (
	"cmp"
)

type NodeType uint8

const (
	TypeLeaf NodeType = iota
	TypeSingleChild
	TypeFull
)

type Node[K cmp.Ordered, V any] struct {
	Key    K           `json:"k,omitempty"`
	Value  V           `json:"v,omitempty"`
	Left   *Node[K, V] `json:"l,omitempty"`
	Right  *Node[K, V] `json:"r,omitempty"`
	Height int         `json:"h,omitempty"`
}

func (node *Node[K, V]) Type() NodeType {
	switch {
	case node.Left != nil && node.Right == nil, node.Left == nil && node.Right != nil:
		return TypeSingleChild
	case node.Left == nil && node.Right == nil:
		return TypeLeaf
	default:
		return TypeFull
	}
}

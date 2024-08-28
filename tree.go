package avltree

import (
	"cmp"
)

// AVLTree is a self-balancing binary search tree.
type AVLTree[K cmp.Ordered, V any] struct {
	root      *Node[K, V]
	size      uint
	threshold int
}

// DefaultThreshold is optimal for fast searching.
const DefaultThreshold = 1

// New returns an AVL tree. The threshold is used for balancing.
// Higher value means faster inserts and deletes but slower searches.
// Recommended value is DefaultThreshold.
func New[K cmp.Ordered, V any](threshold int) *AVLTree[K, V] {
	if threshold < 1 {
		threshold = DefaultThreshold
	}

	return &AVLTree[K, V]{
		threshold: threshold,
	}
}

// RootKey returns the key of the root node.
func (tree *AVLTree[K, V]) RootKey() K {
	return tree.root.Key
}

// Size returns the amounts of nodes in the tree.
func (tree *AVLTree[K, V]) Size() uint {
	return tree.size
}

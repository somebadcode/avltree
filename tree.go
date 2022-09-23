package avltree

// Key is used for the key in the tree and can be anything that supports equality, lesser and greater than.
type Key interface {
	comparable
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// AVLTree is a self-balancing binary search tree.
type AVLTree[K Key, V any] struct {
	root      *Node[K, V]
	size      uint
	threshold int
}

// DefaultThreshold is optimal for fast searching.
const DefaultThreshold = 1

// New returns an AVL tree. The threshold is used for balancing.
// Higher value means faster inserts and deletes but slower searches.
// Recommended value is DefaultThreshold.
func New[K Key, V any](threshold int) *AVLTree[K, V] {
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

func max[T int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64](x, y T) T {
	if x < y {
		return y
	}

	return x
}

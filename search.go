package avltree

import (
	"github.com/somebadcode/avltree/zerovalue"
)

func (tree *AVLTree[K, V]) Search(key K) (V, bool) {
	current := tree.search(key)

	if current == nil {
		return zerovalue.New[V](), false
	}

	return current.Value, true
}

func (tree *AVLTree[K, V]) search(key K) *Node[K, V] {
	for current := tree.root; current != nil; {
		if key == current.Key {
			return current
		}

		if key < current.Key {
			current = current.Left
			continue
		}

		current = current.Right
	}

	return nil
}

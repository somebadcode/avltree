package avltree

func (tree *AVLTree[K, V]) Search(key K) (V, bool) {
	current := tree.search(key)

	if current == nil {
		var zeroV V

		return zeroV, false
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

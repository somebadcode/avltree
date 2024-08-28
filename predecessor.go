package avltree

// InorderPredecessor returns the key and value of the inorder predecessor of the specified key.
func (tree *AVLTree[K, V]) InorderPredecessor(key K) (K, V, bool) {
	// To find the inorder predecessor, we should start at the node which we're getting the successor for.
	node := tree.search(key)

	// The subtree from which we will be searching for the predecessor is the node's left.
	// If the node's left is nil then we need to start searching from the tree's root.
	var subtree *Node[K, V]
	if node != nil && node.Left != nil {
		subtree = node.Left
	} else {
		subtree = tree.root
	}
	// Find the predecessor but if it's nil then it found nothing, return zero key & value, and false.
	node = tree.inorderPredecessor(subtree, key)
	if node == nil {
		var zeroK K
		var zeroV V

		return zeroK, zeroV, false
	}

	// Inorder successor found!
	return node.Key, node.Value, true
}

func (tree *AVLTree[K, V]) inorderPredecessor(subTree *Node[K, V], key K) *Node[K, V] {
	var successor *Node[K, V]

	for current := subTree; current != nil; {
		if current.Key < key {
			successor = current
		}

		if current.Key < key {
			current = current.Right
			continue
		}

		current = current.Left
	}

	return successor
}

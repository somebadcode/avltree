package avltree

// InorderSuccessor returns the key and value of the inorder successor of the specified key.
func (tree *AVLTree[K, V]) InorderSuccessor(key K) (K, V, bool) {
	// To find the inorder successor, we should start at the node which we're getting the successor for.
	node := tree.search(key)

	// The subtree from which we will be searching for the successor is the node's right.
	// If the node's right is nil then we need to start searching from the tree's root.
	var subtree *Node[K, V]
	if node != nil && node.Right != nil {
		subtree = node.Right
	} else {
		subtree = tree.root
	}

	// Find the successor but if it's nil then it found nothing, return zero key & value, and false.
	node = tree.inorderSuccessor(subtree, key)
	if node == nil {
		var zeroK K
		var zeroV V

		return zeroK, zeroV, false
	}

	// Inorder successor found!
	return node.Key, node.Value, true
}

func (tree *AVLTree[K, V]) inorderSuccessor(subTree *Node[K, V], key K) *Node[K, V] {
	var successor *Node[K, V]

	for current := subTree; current != nil; {
		if current.Key > key {
			successor = current
		}

		if current.Key > key {
			current = current.Left
			continue
		}

		current = current.Right
	}

	return successor
}

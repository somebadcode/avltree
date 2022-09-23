package avltree

func (tree *AVLTree[K, V]) Insert(key K, value V) {
	node := &Node[K, V]{
		Key:   key,
		Value: value,
	}

	tree.root = tree.insert(tree.root, node, tree.threshold)
}

func (tree *AVLTree[K, V]) insert(subtree *Node[K, V], node *Node[K, V], threshold int) *Node[K, V] {
	if subtree == nil {
		tree.size++

		return node
	}

	if node.Key == subtree.Key {
		subtree.Value = node.Value

		return subtree
	}

	if node.Key < subtree.Key {
		subtree.Left = tree.insert(subtree.Left, node, threshold)
	} else {
		subtree.Right = tree.insert(subtree.Right, node, threshold)
	}

	subtree.Height = 1 + max(heightOf(subtree.Left), heightOf(subtree.Right))

	return balance(subtree, threshold)
}

package avltree

func (tree *AVLTree[K, V]) Delete(key K) {
	tree.root = tree.delete(tree.root, key)
}

func (tree *AVLTree[K, V]) delete(subtree *Node[K, V], key K) *Node[K, V] {
	if key != subtree.Key {
		// Go down the tree.
		if key < subtree.Key {
			subtree.Left = tree.delete(subtree.Left, key)
		} else {
			subtree.Right = tree.delete(subtree.Right, key)
		}

		return balance(subtree, tree.threshold)
	}

	switch subtree.Type() {
	case TypeLeaf:
		// Leaf nodes can be deleted as is.
		tree.size--

		return nil
	case TypeSingleChild:
		// Single child node.
		if subtree.Left != nil {
			subtree = subtree.Left
			break
		}

		subtree = subtree.Right

	default:
		// Replace current node with inorder successor.
		// We could use inorder predecessor, but then it becomes important that we use the same method for
		// deleting the inorder successor.
		successor := tree.inorderSuccessor(subtree.Right, subtree.Key)

		// Remove the successor from the tree.
		tree.root = tree.delete(tree.root, successor.Key)

		// Replace subtree with successor.
		successor.Left = subtree.Left
		successor.Right = subtree.Right
		successor.Height = subtree.Height
		subtree = successor
	}

	subtree.Height = 1 + max(heightOf(subtree.Left), heightOf(subtree.Right))

	tree.size--

	// Ensure that the tree is balanced before we return.
	return balance(subtree, tree.threshold)
}

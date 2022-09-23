package avltree

// VisitFunc is used when traversing the tree. The function will be called with the key and value.
// The traversal can be stopped by returning false.
type VisitFunc[K Key, V any] func(K, V) bool

// InorderTraversal will do an inorder traversal of the whole tree.
func (tree *AVLTree[K, V]) InorderTraversal(visitFunc VisitFunc[K, V]) {
	inorderTraversal(tree.root, visitFunc)
}

// PreorderTraversal will do a preorder traversal of the whole tree.
func (tree *AVLTree[K, V]) PreorderTraversal(visitFunc VisitFunc[K, V]) {
	preorderTraversal(tree.root, visitFunc)
}

// PostorderTraversal will do a postorder traversal of the whole tree.
func (tree *AVLTree[K, V]) PostorderTraversal(visitFunc VisitFunc[K, V]) {
	postorderTraversal(tree.root, visitFunc)
}

func inorderTraversal[K Key, V any](node *Node[K, V], visitFunc VisitFunc[K, V]) bool {
	if node == nil {
		return true
	}

	if !inorderTraversal(node.Left, visitFunc) {
		return false
	}

	if !visitFunc(node.Key, node.Value) {
		return false
	}

	return inorderTraversal(node.Right, visitFunc)
}

func preorderTraversal[K Key, V any](node *Node[K, V], visitFunc VisitFunc[K, V]) bool {
	if node == nil {
		return true
	}

	if !visitFunc(node.Key, node.Value) {
		return false
	}

	if !preorderTraversal(node.Left, visitFunc) {
		return false
	}

	return preorderTraversal(node.Right, visitFunc)
}

func postorderTraversal[K Key, V any](node *Node[K, V], visitFunc VisitFunc[K, V]) bool {
	if node == nil {
		return true
	}

	if !postorderTraversal(node.Left, visitFunc) {
		return false
	}

	if !postorderTraversal(node.Right, visitFunc) {
		return false
	}

	return visitFunc(node.Key, node.Value)
}

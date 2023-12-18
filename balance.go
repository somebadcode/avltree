package avltree

import (
	"cmp"
)

func balanceOf[K cmp.Ordered, V any](node *Node[K, V]) int {
	if node == nil {
		return 0
	}

	return heightOf(node.Left) - heightOf(node.Right)
}

func heightOf[K cmp.Ordered, V any](node *Node[K, V]) int {
	if node == nil {
		return -1
	}

	return node.Height
}

func balance[K cmp.Ordered, V any](node *Node[K, V], threshold int) *Node[K, V] {
	nodeBalance := balanceOf(node)

	if nodeBalance > threshold {
		// Left heavy node.
		if balanceOf(node.Left) >= 0 {
			node = rotateRight(node)
		} else {
			node = rotateLeftRight(node)
		}

		return node
	}

	if nodeBalance < -threshold {
		if balanceOf(node.Right) <= 0 {
			node = rotateLeft(node)
		} else {
			node = rotateRightLeft(node)
		}
	}

	return node
}

func rotateLeft[K cmp.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	y := x.Right
	x.Right = y.Left
	y.Left = x

	x.Height = 1 + max(heightOf(x.Left), heightOf(x.Right))
	y.Height = 1 + max(heightOf(y.Left), heightOf(y.Right))

	return y
}

func rotateRight[K cmp.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	y := x.Left
	x.Left = y.Right
	y.Right = x

	x.Height = 1 + max(heightOf(x.Left), heightOf(x.Right))
	y.Height = 1 + max(heightOf(y.Left), heightOf(y.Right))

	return y
}

func rotateLeftRight[K cmp.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	node.Left = rotateLeft(node.Left)

	return rotateRight(node)
}

func rotateRightLeft[K cmp.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	node.Right = rotateRight(node.Right)

	return rotateLeft(node)
}

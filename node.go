package avltree

type Node[K Key, V any] struct {
	Key    K           `json:"k,omitempty"`
	Value  V           `json:"v,omitempty"`
	Left   *Node[K, V] `json:"l,omitempty"`
	Right  *Node[K, V] `json:"r,omitempty"`
	Height int         `json:"h,omitempty"`
}

func (node *Node[K, V]) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node *Node[K, V]) HasSingleChild() bool {
	return (node.Left != nil && node.Right == nil) || (node.Left == nil && node.Right != nil)
}

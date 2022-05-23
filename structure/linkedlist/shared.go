package linkedlist

// This node is shared across different implementations.
type Node struct {
	Val  any   // 현재 노드값
	Prev *Node // 이전 노드
	Next *Node // 다음 노드
}

// Create new node.
func NewNode(val any) *Node {
	return &Node{val, nil, nil}
}

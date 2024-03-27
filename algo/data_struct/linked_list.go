package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value, nil}
}
func (nd *Node[T]) SetNextNode(nextNode *Node[T]) {
	nd.next = nextNode
}
func (nd *Node[T]) DeleteNextNode() {
	nd.next = nd.next.next
}
func (nd *Node[T]) GetNextNode() *Node[T] {
	return nd.next
}
func (nd *Node[T]) GetValue() T {
	return nd.value
}
func (nd *Node[T]) PrintCircuit() {
	for nd != nil {
		fmt.Println(nd.value)
		nd = nd.next
	}
}
func main() {
	intMasterNode := NewNode(1)
	int2Node := NewNode(2)
	intMasterNode.SetNextNode(int2Node)
	intMasterNode.PrintCircuit()
	println("----")
	intMasterNode.PrintCircuit()
}

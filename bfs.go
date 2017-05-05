package main

// Tree is our  representation of the tree
// A node is given as the key with its child nodes as value
type Tree map[string][]string

// Stack holds the nodes that are yet to be looked up
// Implements a stack FILO(First in Last Out)
type Stack struct {
	nodeList []string
}

// Pop returns the last element pushed into the Stack
func (s *Stack) Pop() string {
	if len(s.nodeList) < 1 {
		return ""
	}

	var x string
	x, s.nodeList = s.nodeList[len(s.nodeList)-1], s.nodeList[:len(s.nodeList)-1]
	return x
}

// Push adds the given node to the Stack
func (s *Stack) Push(node string) {
	s.nodeList = append(s.nodeList, node)
}

// Len returns the count of nodes to be processed
func (s *Stack) Len() int {
	return len(s.nodeList)
}

// Clear will clear the stack
func (s *Stack) Clear() {
	s.nodeList = nil
}

// FindNode does a BFS on the
func FindNode(tree Tree, rootNode, requiredNode string) (string, error) {
	return "", nil
}

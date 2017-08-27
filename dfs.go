package main

import "errors"

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
func (s *Stack) Push(nodes ...string) {
	s.nodeList = append(s.nodeList, nodes...)
}

// Len returns the count of nodes to be processed
func (s *Stack) Len() int {
	return len(s.nodeList)
}

// Clear will clear the stack
func (s *Stack) Clear() {
	s.nodeList = nil
}

// FindNode does a DFS on the
func FindNode(tree Tree, stack *Stack, requiredNode string) (string, error) {

	currentNode := stack.Pop()
	// found what we are looking for
	if currentNode == requiredNode {
		return requiredNode, nil
	}

	// get the current node child nodes and add them to stack
	nodeList, ok := tree[currentNode]
	if ok {
		stack.Push(nodeList...)
		for stack.Len() != 0 {
			foundNode, err := FindNode(tree, stack, requiredNode)
			if err == nil {
				return foundNode, nil
			}
		}
	}

	return "", errors.New("Node not found")
}

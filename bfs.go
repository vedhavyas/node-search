package main

import "errors"

const MaxQueueSize int = 20

// Queue is the FIFO struct
type Queue struct {
	index    int
	nodeList []string
}

// Push will push the nodes queue
func (q *Queue) Push(nodes ...string) {
	// check if we should drop the visited nodes
	if q.Len()+len(nodes) > MaxQueueSize {
		// drop queue if we have visited all the nodes in the queue
		if q.index >= q.Len() {
			q.nodeList = nil
		} else {
			// copy the to be visited nodes to new queue
			q.nodeList = q.nodeList[q.index:]
		}

		q.index = 0
	}
	q.nodeList = append(q.nodeList, nodes...)
}

// DeQueue will return the first element
func (q *Queue) DeQueue() string {
	if q.Len() < 1 {
		return ""
	}

	deQueued := q.nodeList[q.index]
	q.index++

	return deQueued
}

// Len returns the length of the nodes to be visited
func (q *Queue) Len() int {
	currentSize := len(q.nodeList)
	if q.index >= currentSize {
		return 0
	}

	return currentSize - q.index
}

// FindNode does a DFS on the
func BreadthFirstSearch(tree Tree, queue *Queue, requiredNode string) (string, error) {

	currentNode := queue.DeQueue()
	// found what we are looking for
	if currentNode == requiredNode {
		return requiredNode, nil
	}

	// get the current node child nodes and add them to stack
	nodeList, ok := tree[currentNode]
	if ok {
		queue.Push(nodeList...)
		for queue.Len() != 0 {
			foundNode, err := BreadthFirstSearch(tree, queue, requiredNode)
			if err == nil {
				return foundNode, nil
			}
		}
	}

	return "", errors.New("Node not found")
}

package main

import "testing"

func TestQueue_Push(t *testing.T) {
	q := &Queue{}
	q.Push("a", "b", "c")
	if q.Len() != 3 {
		t.Fatalf("Expected %d but got %d", 3, q.Len())
	}
}

func TestQueue_DeQueue(t *testing.T) {
	q := &Queue{}
	q.Push("a", "b", "c")
	if q.Len() != 3 {
		t.Fatalf("Expected %d but got %d", 3, q.Len())
	}

	deQueuedNode := q.DeQueue()
	if deQueuedNode != "a" {
		t.Fatalf("Expected %s but got %s", "a", deQueuedNode)
	}

	deQueuedNode = q.DeQueue()
	if deQueuedNode != "b" {
		t.Fatalf("Expected %s but got %s", "b", deQueuedNode)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	//          a
	//       /      \
	//      b        c
	//     / \      / \
	//    g   h    d   e
	//            / \
	//           i   j

	tree := make(Tree)
	tree["a"] = []string{"b", "c"}
	tree["b"] = []string{"g", "h"}
	tree["c"] = []string{"d", "e"}
	tree["d"] = []string{"i", "j"}

	testCases := []struct {
		requiredNode string
		result       bool
	}{
		{
			requiredNode: "j",
			result:       true,
		},

		{
			requiredNode: "e",
			result:       true,
		},

		{
			requiredNode: "a",
			result:       true,
		},

		{
			requiredNode: "z",
			result:       false,
		},

		{
			requiredNode: "h",
			result:       true,
		},

		{
			requiredNode: "w",
			result:       false,
		},
	}

	for _, test := range testCases {
		queue := &Queue{}
		queue.Push("a")
		foundNode, err := BreadthFirstSearch(tree, queue, test.requiredNode)
		if err != nil {
			if test.result {
				t.Fatalf("Expected node %s to be found but returned not found!", test.requiredNode)
			}

			continue
		}

		if foundNode != test.requiredNode {
			t.Fatalf("Expected node %s but got %s", test.requiredNode, foundNode)
		}
	}
}

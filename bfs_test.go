package main

import "testing"

func TestStack_Push(t *testing.T) {
	stack := &Stack{}

	if stack.Len() != 0 {
		t.Fatalf("Expected length %d but got %d", 0, stack.Len())
	}

	stack.Push("a")
	if stack.Len() != 1 {
		t.Fatalf("Expected length %d but got %d", 1, stack.Len())
	}

	stack.Push("b")
	if stack.Len() != 2 {
		t.Fatalf("Expected length %d but got %d", 2, stack.Len())
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack{}

	popNode := stack.Pop()
	if popNode != "" {
		t.Fatalf("Expected empty string but got %s", popNode)
	}

	stack.Push("a")
	stack.Push("b")

	popNode = stack.Pop()
	if popNode != "b" {
		t.Fatalf("Expected %s but got %s", "b", popNode)
	}

	stack.Push("c")
	popNode = stack.Pop()
	if popNode != "c" {
		t.Fatalf("Expected %s but got %s", "c", popNode)
	}

	if stack.Len() != 1 {
		t.Fatalf("Expected length %d but got %d", 1, stack.Len())
	}
}

func TestStack_Clear(t *testing.T) {
	stack := Stack{}
	stack.Push("a")
	stack.Push("b")
	if stack.Len() != 2 {
		t.Fatalf("Expected length %d but got %d", 2, stack.Len())
	}

	stack.Clear()
	if stack.Len() != 0 {
		t.Fatalf("Expected length %d but got %d", 0, stack.Len())
	}

	stack.Push("a")
	if stack.Len() != 1 {
		t.Fatalf("Expected length %d but got %d", 1, stack.Len())
	}
}

func TestFindNode(t *testing.T) {
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
		stack := &Stack{}
		stack.Push("a")
		foundNode, err := FindNode(tree, stack, test.requiredNode)
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

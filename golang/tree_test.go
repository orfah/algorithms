package algorithm

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Add(3)
	if tree.Root.Value != 3 {
		t.Error("BST failed Add")
	}

	tree.Add(5)
	if tree.Root.Right.Value != 5 {
		t.Error("BST failed Add")
	}

	tree.Add(1)
	if tree.Root.Left.Value != 1 {
		t.Error("BST failed Add")
	}

	tree.Add(4)
	if tree.Root.Right.Left.Value != 4 {
		t.Error("BST failed 2nd degree Add")
	}

	tree.Add(2)
	if tree.Root.Left.Right.Value != 2 {
		t.Error("BST failed 2nd degree Add")
	}
}

func TestFind(t *testing.T) {
	tree := &BinarySearchTree{}
	if tree.Find(1) != nil {
		t.Error("BST failed Find empty nodes")
	}

	tree = &BinarySearchTree{&TreeNode{2, nil, nil}}
	if tree.Find(2) == nil {
		t.Error("BST failed Find single elt")
	}

	tree.Add(1)
	tree.Add(5)

	if tree.Find(1) == nil || tree.Find(5) == nil {
		t.Error("BST failed Find mult elt")
	}

	if tree.Find(10) != nil {
		t.Error("BST failed Find nonexistent")
	}
}

func TestRemove(t *testing.T) {
	tree := &BinarySearchTree{&TreeNode{1, nil, nil}}
	tree.Remove(1)
	if tree.Root != nil {
		t.Error("BST failed remove root")
	}

	tree = &BinarySearchTree{&TreeNode{1, nil, nil}}
	tree.Add(3)
	tree.Remove(1)
	if tree.Root.Value != 3 || tree.Root.Left != nil || tree.Root.Right != nil {
		t.Error("BST failed remove root")
	}

	tree = &BinarySearchTree{&TreeNode{3, nil, nil}}
	tree.Add(1)
	tree.Add(5)
	tree.Add(4)
	tree.Remove(3)
	if tree.Root.Value != 4 || tree.Root.Left.Value != 1 || tree.Root.Right.Value != 5 {
		t.Error("BST failed remove complex root")
	}

	tree = &BinarySearchTree{&TreeNode{3, nil, nil}}
	tree.Add(2)
	tree.Add(1)
	tree.Remove(3)
	if tree.Root.Value != 2 || tree.Root.Left.Value != 1 {
		t.Error("BST failed remove complex root")
	}
}

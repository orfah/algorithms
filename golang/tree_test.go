package algorithm

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Add(3)
	if tree.Nodes[0].Value != 3 {
		t.Error("BST failed Add")
	}

	tree.Add(4)
	if tree.Nodes[2].Value != 4 {
		t.Error("BST failed Add")
	}

	tree.Add(1)
	if tree.Nodes[1].Value != 1 {
		t.Error("BST failed Add")
	}
}

func TestContains(t *testing.T) {
	tree := &BinarySearchTree{}
	if tree.Contains(1) {
		t.Error("BST failed Contains empty nodes")
	}

	tree = &BinarySearchTree{[]*TreeNode{&TreeNode{1}}}
	if !tree.Contains(1) {
		t.Error("BST failed Contains single elt")
	}

	tree = &BinarySearchTree{
		[]*TreeNode{
			&TreeNode{3}, &TreeNode{1}, &TreeNode{5},
		},
	}
	if !tree.Contains(1) || !tree.Contains(5) {
		t.Error("BST failed Contains mult elt")
	}

	if tree.Contains(10) {
		t.Error("BST failed Contains mult elt")
	}
}

package algorithm

type tree interface {
	Add(interface{})
	Remove(interface{})
}

type BinarySearchTree struct {
	Nodes []*TreeNode
}
type TreeNode struct {
	Value int
}

func (bst *BinarySearchTree) Contains(elt int) bool {
	i := 0
	l := len(bst.Nodes)
	var node *TreeNode
	for i < l {
		node = bst.Nodes[i]
		if node == nil {
			return false
		} else if node.Value == elt {
			return true
		} else if node.Value < elt {
			i = i*2 + 2
		} else {
			i = i*2 + 1
		}
	}

	return false
}

func (bst *BinarySearchTree) Add(elt int) {
	i := 0
	l := len(bst.Nodes)
	var node *TreeNode

	for i < l {
		node = bst.Nodes[i]
		if node == nil {
			break
		} else if elt < node.Value {
			i = i*2 + 1
		} else {
			i = i*2 + 2
		}
	}

	c := cap(bst.Nodes)
	if i >= c {
		nodes := make([]*TreeNode, 4*(c+1))
		copy(nodes, bst.Nodes)
		bst.Nodes = nodes
	}

	bst.Nodes[i] = &TreeNode{elt}
}

package algorithm

type Tree interface {
	Add(interface{}) interface{}
	Remove(interface{})
	Find(interface{}) interface{}
}

type BinarySearchTree struct {
	Root *TreeNode
}
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Returns TreeNode element pointer, or nil
func (bst *BinarySearchTree) Find(elt int) *TreeNode {
	node := bst.Root
	for {
		if node == nil {
			break
		} else if node.Value == elt {
			return node
		} else if node.Value < elt {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	return nil
}

func (bst *BinarySearchTree) Add(elt int) {
	node := bst.Root
	newNode := &TreeNode{elt, nil, nil}

	if node == nil {
		bst.Root = newNode
		return
	}

	for {
		if elt < node.Value {
			if node.Left == nil {
				node.Left = newNode
				return
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = newNode
				return
			}
			node = node.Right
		}
	}
}

func (bst *BinarySearchTree) Remove(elt int) {
	node := bst.Find(elt)
	setRoot := false
	if node == bst.Root {
		setRoot = true
	}

	if node != nil {
		if node.Left != nil && node.Right != nil {
			prevNode := node
			leftMost := node.Right
			for {
				if leftMost.Left == nil {
					break
				}
				prevNode = leftMost
				leftMost = leftMost.Left
			}

			prevNode.Left = leftMost.Right
			node.Value = leftMost.Value
		} else if node.Left == nil {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	if setRoot {
		bst.Root = node
	}
}

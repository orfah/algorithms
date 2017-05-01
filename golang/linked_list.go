package algorithm

type LinkedList interface {
	Add(interface{}) interface{}
	Remove(interface{})
	Contains(interface{}) interface{}
}

type DoublyLinkedList struct {
	Head *DLLNode
	Tail *DLLNode
}

type DLLNode struct {
	Value int
	Prev  *DLLNode
	Next  *DLLNode
}

func (list *DoublyLinkedList) Add(elt int) *DLLNode {
	node := &DLLNode{elt, nil, nil}
	if list.Head == nil {
		list.Head = node
	}
	if list.Tail != nil {
		list.Tail.Next = node
		node.Prev = list.Tail
	}
	list.Tail = node
	return node
}

func (list *DoublyLinkedList) Remove(elt int) {
	node := list.Find(elt)
	if node != nil {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if list.Head == node {
			list.Head = node.Next
		}
		if list.Tail == node {
			list.Tail = node.Prev
		}
		node = nil
	}
}

func (list *DoublyLinkedList) Find(elt int) *DLLNode {
	for node := list.Head; node != nil; {
		if node.Value == elt {
			return node
		}
		node = node.Next
	}
	return nil
}

type SLLNode struct {
	Value int
	Next  *SLLNode
}

type SinglyLinkedList struct {
	Head *SLLNode
	Tail *SLLNode
}

func (list *SinglyLinkedList) Add(elt int) *SLLNode {
	node := &SLLNode{elt, nil}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		list.Tail = node
	}
	return node
}

func (list *SinglyLinkedList) Remove(elt int) *SLLNode {
	return nil
}

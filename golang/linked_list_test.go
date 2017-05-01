package algorithm

import (
	"testing"
)

func TestDLLAdd(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Add(3)
	if dll.Head != dll.Tail || dll.Head.Value != 3 {
		t.Error("Failed DLL Add Head/Tail")
	}

	dll.Add(5)
	if dll.Head.Value != 3 || dll.Tail.Value != 5 {
		t.Error("Failed DLL Add")
	}

	dll.Add(7)
	if dll.Head.Value != 3 || dll.Tail.Value != 7 || dll.Head.Next.Value != 5 {
		t.Error("Failed DLL Add")
	}
}

func TestDLLRemove(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Add(3)
	dll.Remove(3)
	if dll.Head != nil || dll.Tail != nil {
		t.Error("Failed DLL Remove Head/Tail")
	}

	dll.Add(3)
	dll.Add(5)
	dll.Add(7)
	dll.Remove(5)
	if dll.Head.Value != 3 || dll.Head.Next.Value != 7 {
		t.Error("Failed DLL Remove")
	}
}

func TestDLLFind(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Add(3)
	if dll.Find(3) == nil {
		t.Error("Failed DLL Find")
	}
	if dll.Find(5) != nil {
		t.Error("Failed DLL Find")
	}

	dll.Add(5)
	dll.Add(7)
	if dll.Find(5) == nil || dll.Find(7) == nil {
		t.Error("Failed DLL Find")
	}
}

package algorithms

import "testing"

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList(1)
	if ll.Head.Value != 1 {
		t.Errorf("Head set to %d, expected 1", ll.Head.Value)
	}
}

func TestAppend(t *testing.T) {
	ll := NewLinkedList(1)
	ll.add(4)

	current := ll.Head
	if ll.Tail.Value != 4 {
		t.Errorf("Next value is %d, when 4 is expected", ll.Tail.Value)
	}

	if ll.Tail.Prev != current {
		t.Errorf("Prev is not set correctly")
	}
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	ll.add(5)
	ll.remove(3)

	current := ll.Head
	for current.Next != nil {
		if current.Value == 3 {
			t.Error("3 not removed as expected")
		}
		current = current.Next
	}

}

func TestCount(t *testing.T) {
	ll := NewLinkedList(1)

	if ll.count() != 1 {
		t.Errorf("Count should be 1 it is %d", ll.count())
	}

	ll.add(6)

	if ll.count() != 2 {
		t.Errorf("Count should be 2 it is %d", ll.count())
	}
}

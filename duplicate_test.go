package algorithms

import "testing"

func TestDuplicate(t *testing.T) {
	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []string{"hey", "this", "word", "hey", "is", "duplicate"}
		dup := Duplicate(arr)
		if dup != true {
			t.Error("The func did not return true for duplicate")
		}
	})

	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []string{"hey", "this", "word", "is", "duplicate"}
		dup := Duplicate(arr)
		if dup != false {
			t.Error("The func did not return false for duplicate")
		}
	})
}

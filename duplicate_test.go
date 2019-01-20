package algorithms

import "testing"

func TestDuplicateString(t *testing.T) {
	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []string{"hey", "this", "word", "hey", "is", "duplicate"}
		dup := DuplicateString(arr)
		if dup != true {
			t.Error("The func did not return true for duplicate")
		}
	})

	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []string{"hey", "this", "word", "is", "duplicate"}
		dup := DuplicateString(arr)
		if dup != false {
			t.Error("The func did not return false for duplicate")
		}
	})
}

func TestDuplicateNums(t *testing.T) {
	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 1}
		dup := DuplicateNums(arr)
		if dup != true {
			t.Error("The func did not return true for duplicate")
		}
	})

	t.Run("Array has duplicate", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		dup := DuplicateNums(arr)
		if dup != false {
			t.Error("The func did not return false for duplicate")
		}
	})
}

package algorithms

import (
	"reflect"
	"testing"
)

func TestSorts(t *testing.T) {
	sorts := []struct {
		function func(n []int) []int
		name     string
	}{
		{
			BubbleSort,
			"Bubble",
		},
	}

	t.Run("Test nums are sorted corectly", func(t *testing.T) {
		for _, sort := range sorts {
			actual := sort.function([]int{4, 2, 7, 1, 3})
			expected := []int{1, 2, 3, 4, 7}
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Actual %v Expected: %v", actual, expected)
			}
		}

	})
}

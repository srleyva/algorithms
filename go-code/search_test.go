package algorithms

import (
	"testing"
)

func TestSearches(t *testing.T) {
	array := makeRange(0, 5000000)
	searches := []struct {
		function func(n []int, m int) int
		name     string
	}{
		{
			LinearSearch,
			"Linear",
		},
		{
			BinarySearch,
			"Binary",
		},
	}

	t.Run("Test num is in array", func(t *testing.T) {
		num := 40
		for _, search := range searches {
			index := search.function(array, num)
			if index != 40 {
				t.Errorf("Wrong index returned: %d for %s", index, search.name)
			}
		}

	})

	t.Run("Test num isn't in array", func(t *testing.T) {
		num := -100
		for _, search := range searches {
			index := search.function(array, num)
			if index != -1 {
				t.Errorf("Wrong index returned: %d for %s", index, search.name)
			}
		}

	})
}

func BenchmarkSearches(b *testing.B) {
	// Worst Case
	array := makeRange(0, 5000000)
	num := -100

	searches := []struct {
		function func(n []int, m int) int
		name     string
	}{
		{
			LinearSearch,
			"Linear",
		},
		{
			BinarySearch,
			"Binary",
		},
	}

	for _, search := range searches {
		b.Run(search.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				search.function(array, num)
			}
		})
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

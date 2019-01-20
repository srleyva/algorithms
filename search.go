package algorithms

// LinearSearch for comparison
func LinearSearch(n []int, i int) int {
	for j := 0; j < len(n); j++ {
		if n[j] == i {
			return j
		}
	}
	return -1
}

// BinarySearch searches for i in n using Binary Search
func BinarySearch(n []int, i int) int {
	lowerBound := 0
	upperBound := len(n) - 1
	for lowerBound <= upperBound {
		mid := (lowerBound + upperBound) / 2
		if n[mid] == i {
			return mid
		} else if n[mid] > i {
			upperBound = mid - 1
		} else if n[mid] < i {
			lowerBound = mid + 1
		}
	}
	return -1
}

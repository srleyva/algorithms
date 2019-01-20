package algorithms

func Duplicate(arr []string) bool {
	for j, char := range arr {
		for i, comp := range arr {
			if j != i && comp == char {
				return true
			}
		}
	}
	return false
}

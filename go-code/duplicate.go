package algorithms

func DuplicateString(arr []string) bool {
	seen := make(map[string]bool)
	for _, char := range arr {
		if seen[char] == true {
			return true
		}
		seen[char] = true
	}
	return false
}

func DuplicateNums(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[num] == true {
			return true
		}
		seen[num] = true
	}
	return false
}

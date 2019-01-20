package algorithms

// BubbleSort is the sort implementation that takes the biggest number and bubbles it to the top
func BubbleSort(nums []int) []int {
	unsorted := (len(nums) - 1)
	for unsorted != 0 {
		for i := range nums {
			if i < (unsorted) && nums[i] > nums[i+1] {
				buffer := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = buffer
			}
		}
		unsorted = unsorted - 1
	}
	return nums
}

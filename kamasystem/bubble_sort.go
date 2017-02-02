package kamasystem

/* BubbleSort algorithm using the simplicity and power of GOlang */

func BubbleSort(nums []int) []int {
	sorted := nums[:]
	n := len(sorted)
	for j:= n; j>1; j-- {
		for i := 0; i < n-1; i++ {
			k := i+1
			if (sorted[i] > sorted[k]) {
				sorted[k], sorted[i] = sorted[i], sorted[k]
			}
		}
		sorted = sorted[:j-1]
		n = len(sorted)
	}
	return nums
}

func BubbleSortDesc(nums []int) []int {
	sorted := nums[:]
	n := len(sorted)
	for j:= n; j>1; j-- {
		for i := 0; i < n-1; i++ {
			k := i+1
			if (sorted[i] < sorted[k]) {
				sorted[k], sorted[i] = sorted[i], sorted[k]
			}
		}
		sorted = sorted[:j-1]
		n = len(sorted)
	}
	return nums
}
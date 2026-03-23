package easy

func sortColors(nums []int) {
	count := [3]int{0, 0, 0}
	for _, n := range nums {
		count[n]++
	}
	index := 0
	for d, c := range count {
		for i := 0; i < c; i++ {
			nums[index] = d
			index++
		}
	}
}

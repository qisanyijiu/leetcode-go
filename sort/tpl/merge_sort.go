package tpl

/**
归并排序
*/

func MergeSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	mergeSort(nums, 0, length - 1)
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	var temp []int
	i, j, k := start, 0, mid+1
	for ; i <= mid && k <= end; j++ {
		if nums[i] < nums[k] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[k])
			k++
		}
	}
	if i > mid {
		temp = append(temp, nums[k:end+1]...)
	} else {
		temp = append(temp, nums[i:mid+1]...)
	}
	copy(nums[start:end+1], temp)
}

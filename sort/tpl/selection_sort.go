package tpl

/**
选择排序

首先, 在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。
 */

func SelectionSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i ++ {
		minIndex := i
		for j := i + 1; j < length; j ++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

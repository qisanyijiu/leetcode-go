package tpl

/**
快速排序

快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

在数组中，选择一个元素作为“基准”(pivot),一般选择第一个元素作为基准元素。设置两个游标i和j，初始时i指向数组首元素，j指向尾元素。
从数组最右侧向前扫描，遇到小于基准值的元素停止扫描，将两者交换，然后从数组左侧开始扫描，遇到大于基准值的元素停止扫描，同样将两者交换。
i==j时分区完成，否则转2。
*/

//func QuickSort(nums []int) {
//	start, end := 0, len(nums)-1
//	if start >= end {
//		return
//	}
//	quickSortHelper(nums, start, end)
//}
//
//func partition(nums []int, left, right int) int {
//	if len(nums) < 2 {
//		return right
//	}
//	pivot := nums[right]
//	i := left
//	for j := left; j < right; j++ {
//		if nums[j] < pivot {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//		}
//	}
//	nums[i], nums[right] = nums[right], nums[i]
//	return i
//}
//
//func quickSortHelper(nums []int, left, right int) {
//	if left >= right {
//		return
//	}
//	pivotIndex := partition(nums, left, right)
//	quickSortHelper(nums, left, pivotIndex-1)
//	quickSortHelper(nums, pivotIndex+1, right)
//}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, left, right int) int {
	if len(arr) < 2 {
		return right
	}
	pivot := arr[right] // 参考
	i := left
	for j := left; j < right; j++ {
		// 比参考小
		if arr[j] < pivot {
			// 左移
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// i 左边的都比pivot小，i 右边的逗比pivote大
	// i 和 right调换就行
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	pivoteIndex := partition(arr, left, right)
	quickSort(arr, left, pivoteIndex-1)
	quickSort(arr, pivoteIndex+1, right)
}

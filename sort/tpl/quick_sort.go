package tpl

/**
快速排序

快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

在数组中，选择一个元素作为“基准”(pivot),一般选择第一个元素作为基准元素。设置两个游标i和j，初始时i指向数组首元素，j指向尾元素。
从数组最右侧向前扫描，遇到小于基准值的元素停止扫描，将两者交换，然后从数组左侧开始扫描，遇到大于基准值的元素停止扫描，同样将两者交换。
i==j时分区完成，否则转2。
*/

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, left, right int) int {
	if left >= right {
		return right
	}
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(arr, left, right)
	quickSort(arr, left, pivotIndex-1)
	quickSort(arr, pivotIndex+1, right)
}

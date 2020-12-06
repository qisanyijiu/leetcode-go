package tpl

/**
快速排序

快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

在数组中，选择一个元素作为“基准”(pivot),一般选择第一个元素作为基准元素。设置两个游标i和j，初始时i指向数组首元素，j指向尾元素。
从数组最右侧向前扫描，遇到小于基准值的元素停止扫描，将两者交换，然后从数组左侧开始扫描，遇到大于基准值的元素停止扫描，同样将两者交换。
i==j时分区完成，否则转2。
*/

func QuickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	quickSort(nums, 0, length - 1)
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := partition(nums, start, end)
	quickSort(nums, start, mid)
	quickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	temp := nums[start]
	low := start
	height := end

	for low < height {

		for low < height && temp < nums[height] {
			height--
		}

		if low < height {
			nums[low] = nums[height]
		}

		for low < height && temp > nums[low] {
			low++
		}

		if low < height {
			nums[height] = nums[low]
		}
	}
	nums[low] = temp
	return low
}

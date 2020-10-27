package tpl

/**
插入排序

将待排序的数组划分为局部有序子数组subSorted和无序子数组subUnSorted
每次排序时从subUnSorted中挑出第一个元素
从后向前将其与subSorted各元素比较大小
按照大小插入合适的位置
插入完成后将此元素从subUnSorted中移除
重复这个过程直至subUnSorted中没有元素
总之就时从后向前
一边比较一边移动
*/

func InsertionSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 1 ; i < length; i ++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j -- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}

}

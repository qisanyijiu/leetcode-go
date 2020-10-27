package tpl

/**
希尔排序

希尔排序基于插入排序发展而来。希尔排序的思想基于两个原因：

当数据项数量不多的时候，插入排序可以很好的完成工作。
当数据项基本有序的时候，插入排序具有很高的效率。
基于以上的两个原因就有了希尔排序的步骤：

a.将待排序序列依据步长(增量)划分为若干组，对每组分别进行插入排序。初始时，step=len/2，此时的增量最大，因此每个分组内数据项个数相对较少，插入排序可以很好的完成排序工作（对应1）。

b.以上只是完成了一次排序，更新步长step=step/2,每个分组内数据项个数相对增加，不过由于已经进行了一次排序，数据项基本有序，此时插入排序具有更好的排序效率(对应2)。直至增量为1时，此时的排序就是对这个序列使用插入排序，此次排序完成就表明排序已经完成。
*/

func ShellSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	step := length / 2
	for ; step > 0; step = step / 2 {
		for i := step; i < length; i++ {
			j := i - step
			temp := nums[i]
			for j >= 0 && temp < nums[j] {
				nums[j+step] = nums[j]
				j = j - step
			}
			nums[j+step] = temp
		}
	}
}

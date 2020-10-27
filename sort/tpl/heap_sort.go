package tpl

/**
堆排序

堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。
不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成
*/

func HeapSort(nums []int) {
	for i := len(nums)/2; i>=0; i-- {
		initHead(nums,i,len(nums))
	}

	for i := len(nums)-1;i >0; i--{
		nums[0],nums[i] = nums[i],nums[0]
		initHead(nums,0,i)
	}
}

func initHead(nums []int,parent,len int){
	temp := nums[parent]
	child := 2*parent+1

	for child < len {
		if child+1 < len && nums[child] < nums[child+1] {
			child++
		}

		if child < len && nums[child] <= temp {
			break
		}

		nums[parent] = nums[child]

		parent = child
		child = child*2+1
	}

	nums[parent] = temp
}
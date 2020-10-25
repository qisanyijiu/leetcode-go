package easy

type NumArray struct {
	Value []int
}

func Constructor(nums []int) NumArray {
	// 一维前缀和
	numa := NumArray{[]int{0}}  // 浪费第一个空间
	for i, v := range nums {
		numa.Value = append(numa.Value, v + numa.Value[i])
	}
	return numa
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Value[j+1] - this.Value[i]
}
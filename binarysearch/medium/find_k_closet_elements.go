package medium

import (
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	left, right, mid := 0, l-1, 0
	// 二分查找确定大概位置
	for left < right {
		mid = (left+right)/2
		if arr[mid] == x {
			break
		}
		if arr[mid] > x {
			right = mid
			continue
		}
		left = mid +1
	}
	bottom := int(math.Max(0, float64(mid-k)))
	top := int(math.Min(float64(l-1), float64(mid+k)))
	list := arr[bottom:top+1] // 截取数组
	sort.Slice(list, func(i, j int) bool {
		if math.Abs(float64(list[i]-x)) < math.Abs(float64(list[j]-x)) {
			return true
		}
		if math.Abs(float64(list[i]-x)) == math.Abs(float64(list[j]-x)) &&
			list[i] < list[j] {
			return true
		}
		return false
	})
	list = list[:k]
	sort.Ints(list)
	return list
}
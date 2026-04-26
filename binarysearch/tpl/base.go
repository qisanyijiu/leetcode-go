package tpl

func BinarySearch(a []int, x int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)>>1
		if a[mid] < x {
			left = mid + 1
		} else if a[mid] > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

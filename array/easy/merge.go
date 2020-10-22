package easy

// 原地交换，时间复杂度O(n), 空间复杂度O(1)
func Merge(nums1 []int, m int, nums2 []int, n int)  {
	for m >= 1 && n >= 1 {
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
	for n >= 1 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
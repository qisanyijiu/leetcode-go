package medium

import "log"

//func permute(nums []int) [][]int {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	ans := make([][]int, 0)
//	backtrack(nums, nil, &ans)
//
//	return ans
//}
//
//func backtrack(nums []int, prev []int, ans *[][]int) {
//	log.Println(nums, prev)
//	if len(nums) == 0 {
//		*ans = append(*ans, append([]int{}, prev...))
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		backtrack(
//			append(append([]int{}, nums[0:i]...), nums[i+1:]...),
//			append(prev, nums[i]),
//			ans)
//	}
//}
func permute(nums []int) [][]int{
	if len(nums) == 0{
		return nil
	}
	ans := make([][]int, 0)
	backtrack(nums, []int{}, &ans)
	return ans
}

func backtrack(nums []int, prev []int, ans *[][]int){
	log.Println(nums, prev)
	if len(nums) == 0{
		*ans = append(*ans, prev)
		return
	}
	for i := 0; i < len(nums); i ++ {
		backtrack(
			append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ans)
	}
}
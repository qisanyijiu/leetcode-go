package medium

import "sort"

func threeSum(nums []int) [][]int{
	result := [][]int{}
	dict := map[int]int{}
	for _, num := range nums{
		if _, ok := dict[num]; ok {
			dict[num] += 1
		}else{
			dict[num] = 1
		}
	}
	nums = []int{}
	for k, _ := range dict{
		nums = append(nums, k)
	}
	sort.Ints(nums)
	for i, n1 := range nums{
		for _, n2 := range nums[i:]{
			if n1 == n2 && dict[n1] < 2{
				continue
			}
			n3 := 0 - n1 - n2
			if cnt, ok := dict[n3]; ok {
				if n1 == n3{
					cnt -= 1
				}
				if n2 == n3{
					cnt -= 1
				}
				if cnt > 0 {
					record := []int{n1, n2, n3}
					sort.Ints(record)
					if !inThreeSums(record, result){
						result = append(result, record)
					}
				}
			}
		}
	}
	return result
}

func inThreeSums(item []int,res [][]int) bool {
	for _, row := range res{
		if item[0] == row[0] && item[1] == row[1]&& item[2] == row[2]{
			return true
		}
	}
	return false
}
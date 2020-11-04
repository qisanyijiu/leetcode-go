package medium

import "leetcode/utils"

func MaxScore(cardPoints []int, k int) int {
	leftSum := 0
	rightSum := 0
	for i := 0; i < k; i++ {
		leftSum += cardPoints[i]
	}
	max := leftSum + rightSum
	for i := 0; i < k; i++ {
		leftSum -= cardPoints[k-1-i]
		rightSum += cardPoints[len(cardPoints) - 1 -i]
		tmp := leftSum + rightSum
		max = utils.MaxInt(max, tmp)
	}
	return max
}
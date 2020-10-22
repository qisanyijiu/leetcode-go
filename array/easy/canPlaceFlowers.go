package easy

func CanPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i, v := range flowerbed{
		if v == 0 && (i==0 || flowerbed[i-1]==0) && (i==len(flowerbed)-1 || flowerbed[i+1]==0) {
			flowerbed[i] ++
			cnt ++
		}

		if cnt >= n {
			return true
		}
	}
	return false
}

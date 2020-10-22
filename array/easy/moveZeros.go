package easy


func MoveZeroes(nums []int)  {
	lNonZ := 0;             //当前非0零序列长度

	for i:= range(nums) {
		if (nums[i] != 0) {  //遇到非0数，将其添加到非0序列队尾，并更新长度
			if(i != lNonZ) { //当前数不在非0序列尾，说明含有0序列，将其与0序列头交换
				nums[lNonZ] = nums[i];
				nums[i] = 0;
			}
			lNonZ++;
		}
	}
}
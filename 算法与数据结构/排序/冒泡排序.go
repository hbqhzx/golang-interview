package 排序

func bubbleSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums)-1; i++ { //控制已经ready的数量
		flag := true
		for j := 0; j < len(nums)-i-1; j++ { //从0->len()-ready数量
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}

		if flag {
			return
		}
	}
}

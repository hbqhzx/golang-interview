package 排序

func quickSort(nums []int, left, right int) {
	if len(nums) == 0 || left > right {
		return
	}

	i, j := left, right
	for i < j {
		//当我们把基准值设置在左边时，为什么要从右边先开始找？
		//一定要先从右边找j ！！！
		for nums[j] >= nums[left] && j > i {
			j--
		}

		for nums[i] <= nums[left] && i < j {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}

	nums[left], nums[i] = nums[i], nums[left]

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

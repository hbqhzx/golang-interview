package main

//双指针

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] == val {
			fast++
			continue
		}
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return slow
}

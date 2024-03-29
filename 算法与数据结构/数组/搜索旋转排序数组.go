package main

//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
//你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
//示例 1：
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//示例 3：
//
//输入：nums = [1], target = 0
//输出：-1
//

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2

	if nums[middle] == target {
		return middle
	}
	//1分为2，分成有序和无序两个部分， 看有序部分的左右边界取值，无序部分直接取另一边界
	if nums[left] <= nums[middle] {
		if nums[left] <= target && target < nums[middle] { //注意这里的起始=号,即需要关注有序区间是否包括左等于号。 如果目标值=left，是可以归纳到【left, mid-1】, 如果目标值=middle，上面已经处理过了，所以不用考虑等于，
			return binarySearch(nums, left, middle-1, target)
		} else {
			return binarySearch(nums, middle+1, right, target)
		}

		//nums[left] > nums[middle]
	} else {
		if nums[left] > target && target > nums[middle] { //这里有序的是右区间，所以nums[left]不能等于target，不然就是右区间+left了
			return binarySearch(nums, middle+1, right, target)

		} else {
			return binarySearch(nums, left, middle-1, target)

		}
	}
}

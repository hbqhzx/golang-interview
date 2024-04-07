package main

import (
	"fmt"
	"math"
)

//滑动窗口法

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	window_start := 0
	minSubLen := math.MaxInt

	for window_end := 0; window_end < len(nums); window_end++ {
		sum = sum + nums[window_end]

		for sum >= target {
			window_len := window_end - window_start + 1
			if window_len < minSubLen {
				minSubLen = window_len
			}

			sum = sum - nums[window_start]
			window_start++
		}
	}

	if minSubLen == math.MaxInt {
		return 0
	}

	return minSubLen
}

func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(res)
}

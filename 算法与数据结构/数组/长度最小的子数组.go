package main

import "fmt"

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
	rlen := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		rlen++
		if sum < target {
			continue
		} else {
			if res < rlen || res == 0 {
				res = rlen
			}
			sum = sum - nums[i-rlen+1]
			rlen--
		}
	}

	return res

}

func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(res)
}

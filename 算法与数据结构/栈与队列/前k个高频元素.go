package main

import "sort"

// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:
//
// 输入: nums = [1], k = 1
// 输出: [1]
func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}

	countNumsMap := map[int][]int{}
	for num, count := range numMap {
		countNumsMap[count] = append(countNumsMap[count], num)
	}

	counts := []int{}
	for count := range countNumsMap {
		counts = append(counts, count)
	}

	sort.Ints(counts)
	res := []int{}

	for i := len(counts) - 1; i >= 0 && len(res) < k; {
		cur := countNumsMap[counts[i]]
		for j := 0; len(res) < k && j < len(cur); j++ {
			res = append(res, cur[j])
		}
		i--
	}
	return res

}

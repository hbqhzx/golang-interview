package main

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]

var results3 [][]int
var path3 []int

func permute(nums []int) [][]int {
	results3 = [][]int{}
	path3 = []int{}
	backtracing2(nums)
	return results3
}

func backtracing2(nums []int) {
	if len(path3) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path3)
		results3 = append(results3, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if isContain(path3, nums[i]) {
			continue
		}
		path3 = append(path3, nums[i])
		backtracing2(nums)
		path3 = path3[:len(path3)-1]
	}
}

func isContain(path []int, value int) bool {
	for _, p := range path {
		if p == value {
			return true
		}
	}
	return false
}

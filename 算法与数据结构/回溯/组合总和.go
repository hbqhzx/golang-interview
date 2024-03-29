package main

import (
	"fmt"
)

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为： [ [7], [2,2,3] ]
//示例 2：
//
//输入：candidates = [2,3,5], target = 8,
//所求解集为： [ [2,2,2,2], [2,3,3], [3,5] ]

var paths2 []int
var allPath [][]int

func combinationSum(candidates []int, target int) [][]int {
	paths2 = []int{}
	allPath = [][]int{}

	dfs_v1(candidates, target, 0)
	return allPath

}

func dfs_v1(candidates []int, target int, start int) {
	s := sum(paths2)
	if s == target {
		temp := make([]int, len(paths2)) ////注意，这里一定要初始化长度，不然copy的结果是空
		copy(temp, paths2)
		allPath = append(allPath, temp)
		return
	}
	if s > target {
		return
	}

	// 每次遍历都只需要遍历当前数字及后续的数字就可以，第一次写这里是0索引开始，会重复遍历，即 2222，2232,2322等问题
	for i := start; i < len(candidates); i++ {
		paths2 = append(paths2, candidates[i])
		dfs_v1(candidates, target, i)
		paths2 = paths2[:len(paths2)-1]
	}
}

func sum(value []int) int {
	s := 0
	for _, v := range value {
		s = s + v
	}
	return s
}

func main() {
	c := []int{2, 3, 6, 7}

	fmt.Println(combinationSum(c, 7))
}

package main

import (
	"fmt"
	"strconv"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

var lefferMap = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

var path string
var paths []string

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	path = ""
	paths = []string{}
	if digits == "" {
		return []string{}
	}
	ds := []int{}
	for _, d := range digits {
		i, _ := strconv.Atoi(string(d))
		ds = append(ds, i)
	}
	dfs(ds, 0)
	return paths
}

func dfs(digits []int, index int) {
	if len(path) == len(digits) {
		paths = append(paths, path)
		return
	}

	value := digits[index]
	letters := lefferMap[value]
	for _, l := range letters {
		path = path + l
		dfs(digits, index+1)
		path = path[:len(path)-1]
	}
}

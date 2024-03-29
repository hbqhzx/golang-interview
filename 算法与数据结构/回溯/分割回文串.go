package main

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
//回文串
// 。返回 s 所有可能的分割方案。
//
//
//
//示例 1：
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//示例 2：
//
//输入：s = "a"
//输出：[["a"]]

var res [][]string
var path2 []string

func partition(s string) [][]string {
	res = [][]string{}
	path2 = []string{}
	backtracing(s, 0)
	return res
}

func backtracing(s string, startIndex int) {
	if startIndex == len(s) { //{ // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		temp := make([]string, len(path2)) //注意，这里一定要初始化长度，不然copy的结果是空
		copy(temp, path2)
		res = append(res, temp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		//如果是回文串
		v := s[startIndex : i+1] // i+1的边界不算
		if isCircle(v) {
			path2 = append(path2, v)
			backtracing(s, i+1)
			path2 = path2[:len(path2)-1]
		}
	}
}

func isCircle(v string) bool {

	for i := 0; i < len(v)/2; i++ {
		if v[i] != v[len(v)-1-i] {
			return false
		}
	}
	return true
}

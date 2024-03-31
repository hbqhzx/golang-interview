package main

import "fmt"

// 滑动窗口

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长
// 子串
// 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func lengthOfLongestSubstring(s string) int {
	right, maxSubLen := 0, 0
	subMap := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		//移动左指针，删除上一个左数
		if i != 0 {
			delete(subMap, s[i-1])
		}
		//递增右边界
		for right < len(s) && !subMap[s[right]] {

			subMap[s[right]] = true
			right++
		}

		if len(subMap) > maxSubLen {
			maxSubLen = len(subMap)
		}
	}

	return maxSubLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdfgsfs"))
}

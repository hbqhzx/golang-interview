package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range []byte(s) {
		if left, ok := m[v]; ok {
			if len(stack) == 0 {
				return false
			}
			//出栈一个
			last := stack[len(stack)-1]
			if last != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

package main

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//示例: 输入: n = 4, k = 2 输出: [ [2,4], [3,4], [2,3], [1,2], [1,3], [1,4], ]

// 数组写在外面的好处是，不用担心指针问题，但每次触发时要记得初始化
var (
	result  []int
	results [][]int
)

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}
	results, result = make([][]int, 0), make([]int, 0)
	backtracing_combine(1, n, k)
	return results

}

func backtracing_combine(index int, n, k int) {
	if len(result) == k {
		combination := make([]int, k) ////注意，这里一定要初始化长度，不然copy的结果是空
		copy(combination, result)
		results = append(results, combination)
		return
	}

	if len(result)+(n+1-index) < k {
		return
	}

	for i := index; i <= n; i++ {
		result = append(result, i)
		backtracing_combine(i+1, n, k)
		result = result[:len(result)-1]
	}
}

//该代码中存在一个重大的错误。在使用切片result作为结果添加到结果集results中时，
//你没有对result进行深拷贝。而Go语言的切片是一个引用类型，当你修改result的时候，也会影响到添加到results的切片。
//例如，当你从result中移除一个元素时，也会在results中复制的那个元素被移除。

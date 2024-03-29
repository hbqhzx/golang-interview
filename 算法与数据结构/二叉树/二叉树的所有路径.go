package binaryTree

import (
	"strconv"
	"strings"
)

// 第二种解法
var path []string
var paths [][]string

func binaryTreePaths(root *TreeNode) []string {
	path = []string{}
	paths = [][]string{}
	back(root)
	res := []string{}
	for _, p := range paths {
		res = append(res, strings.Join(p, "->"))
	}
	return res
}

func back(node *TreeNode) {
	if node == nil {
		return
	}
	path = append(path, strconv.Itoa(node.Val))

	if node.Right == nil && node.Left == nil {
		tmp := make([]string, len(path))
		copy(tmp, path)
		paths = append(paths, tmp)
		return
	}

	if node.Left != nil {
		back(node.Left)
		path = path[:len(path)-1]
	}

	if node.Right != nil {
		back(node.Right)
		path = path[:len(path)-1]
	}

}

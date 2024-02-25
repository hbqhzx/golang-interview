package binaryTree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	path := ""
	paths := []string{}
	dfs(root, path, &paths)
	return paths
}

func dfs(tree *TreeNode, path string, paths *[]string) {
	if tree == nil {
		return
	}
	if path != "" {
		path = fmt.Sprintf("%s->%d", path, tree.Val)
	} else {
		path = fmt.Sprintf("%d", tree.Val)
	}

	if tree.Left == nil && tree.Right == nil {
		*paths = append(*paths, path)
		return
	}

	dfs(tree.Left, path, paths)
	dfs(tree.Right, path, paths)
}

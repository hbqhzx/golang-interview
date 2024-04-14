package binaryTree

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。

var paths [][]int
var path []int

func pathSum(root *TreeNode, targetSum int) [][]int {
	paths = [][]int{}
	path = []int{}
	backtracing(root, targetSum, 0)
	return paths
}

func backtracing(node *TreeNode, targetSum int, sum int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	sum = sum + node.Val
	if node.Left == nil && node.Right == nil {
		if sum == targetSum {
			temp := make([]int, len(path))
			copy(temp, path)
			paths = append(paths, temp)
		}
	}

	backtracing(node.Left, targetSum, sum)
	backtracing(node.Right, targetSum, sum)
	path = path[:len(path)-1]
}

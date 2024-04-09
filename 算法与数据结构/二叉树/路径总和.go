package binaryTree

//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
//
//叶子节点 是指没有子节点的节点。

func hasPathSum(root *TreeNode, targetSum int) bool {
	return backtrace(root, targetSum, 0)
}

func backtrace(node *TreeNode, target int, sum int) bool {
	if node == nil {
		return false
	}

	sum = sum + node.Val
	if node.Right == nil && node.Left == nil {
		if sum == target {
			return true
		}
		return false
	}

	return backtrace(node.Left, target, sum) || backtrace(node.Right, target, sum)

}

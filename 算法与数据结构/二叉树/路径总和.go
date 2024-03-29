package binaryTree

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

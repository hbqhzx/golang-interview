package binaryTree

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)
	return max(leftDep, rightDep) + 1
}

// 二叉树最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1

}

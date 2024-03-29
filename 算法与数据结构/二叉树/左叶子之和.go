package binaryTree

// 在父节点之上判断
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftval := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftval = root.Left.Val
	}

	return leftval + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

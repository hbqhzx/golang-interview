package binaryTree

// 判断二叉树是否是平衡二叉树
// 可以是空树。  假如不是空树，任何一个结点的左子树与右子树都是平衡二叉树，并且高度之差的绝对值不超过 1。
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)
	if abs(leftDep, rightDep) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

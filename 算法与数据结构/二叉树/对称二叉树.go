package binaryTree

// 对称二叉树--二叉树是否对称
//
//	1
//
// 2   2
// 3 4 4 3

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)

}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

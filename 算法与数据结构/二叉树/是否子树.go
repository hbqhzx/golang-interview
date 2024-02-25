package binaryTree

// 一个树是否是另外一个的子树
func isSubtree(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil {
		return false
	}
	return check(tree1, tree2) || isSubtree(tree1.Left, tree2) || isSubtree(tree1.Right, tree2)

}

// 检查两个树是否相同
func check(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return check(tree1.Left, tree2.Left) && check(tree1.Right, tree2.Right)
}

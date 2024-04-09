package binaryTree

//给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
//二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树

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

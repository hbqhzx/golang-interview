package binaryTree

//扩展问题
// 二叉树的节点个数，遍历
// 二叉树叶子结点个数，层序遍历，无子节点的就是

// 二叉树中的节点个数【递归】
func numOfTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNum := numOfTreeNode(root.Left)
	rightNum := numOfTreeNode(root.Right)
	return leftNum + rightNum + 1
}

// 二叉树中叶子节点个数【递归】
func numOfNoChildNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return numOfNoChildNode(root.Left) + numOfNoChildNode(root.Right)
}

// 求二叉树第k层节点个数
func numOfthKNode(root *TreeNode, k int) int {
	if root == nil || k <= 0 {
		return 0
	}

	if k == 1 {
		return 1
	}

	return numOfthKNode(root.Right, k-1) + numOfthKNode(root.Left, k-1)
}

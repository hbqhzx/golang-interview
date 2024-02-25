package binaryTree

// 可以从root开始找，判断左子树或者右子树是否有p或者q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left

	} else {
		return root
	}

}

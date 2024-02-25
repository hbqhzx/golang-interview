package binaryTree

import "math"

//二叉搜索树是一个有序树：
//
//若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
//若它的右子树不空，则右子树上所有结点的值

//验证是否是二叉搜索树
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//不能单纯的比较左节点小于中间节点，右节点大于中间节点就完事了。 右子树可能出现左节点小于当前节点，但是也小于根节点

// 中序遍历时，判断当前节点是否大于中序遍历的前一个节点，如果大于，说明满足 BST，继续遍历；否则直接返回 false。
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{}
	preNodeVal := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tree := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tree.Val <= preNodeVal {
			return false
		}
		preNodeVal = tree.Val
		root = tree.Right
	}
	return true
}

// 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

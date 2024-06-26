package binaryTree

import "math"

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

func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
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

// 二叉搜索树的最小绝对差，中序遍历所有值，然后比较差值
func getMinimumDifference(root *TreeNode) int {
	values := inorderTraversal(root)
	diff := math.MaxInt64
	for i := 0; i < len(values)-1; i++ {
		cd := values[i+1] - values[i]
		if cd < diff {
			diff = cd
		}
	}
	return diff
}

// 中序递归逻辑中比较差值
func getMinimumDifferenceV2(root *TreeNode) int {
	stack := []*TreeNode{}
	res := math.MaxInt64
	var preNode *TreeNode = nil
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if preNode != nil {
				res = min(res, root.Val-preNode.Val)
			}
			preNode = root
			root = root.Right
		}
	}
	return res
}

// 二叉搜索树中的众数
// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//
// 如果树中有不止一个众数，可以按 任意顺序 返回。
//
// 假定 BST 满足如下定义：
//
// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树

var maxCount, curCount, preVal int
var res []int

func findMode(root *TreeNode) []int {
	maxCount, curCount, preVal = 0, 0, 0
	res = []int{}
	if root == nil {
		return []int{}
	}
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)

	if root.Val != preVal {
		curCount = 1
	} else {
		curCount++
	}

	if curCount > maxCount {
		maxCount = curCount
		res = []int{root.Val}
	} else if curCount == maxCount {
		res = append(res, root.Val)
	}
	preVal = root.Val

	dfs(root.Right)
}

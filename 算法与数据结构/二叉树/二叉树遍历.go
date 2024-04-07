package binaryTree

// 二叉树前序遍历
// 递归
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 前序  中左右---， 【推荐】入栈顺序，先右后左,画个U作为栈，感受下过程
func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tree := stack[len(stack)-1]
		res = append(res, tree.Val)
		stack = stack[:len(stack)-1]
		if tree.Right != nil {
			stack = append(stack, tree.Right)
		}
		if tree.Left != nil {
			stack = append(stack, tree.Left)
		}
	}
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil { //访问到最底层
			stack = append(stack, root) //将访问的节点放进栈
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right

	}
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tree := stack[len(stack)-1]
		res = append(res, tree.Val)
		stack = stack[:len(stack)-1]
		if tree.Left != nil { //// 相对于前序遍历，这更改一下入栈顺序 （空节点不入栈）
			stack = append(stack, tree.Left)
		}
		if tree.Right != nil {
			stack = append(stack, tree.Right)
		}
	}

	newRes := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		newRes = append(newRes, res[i])
	}
	return newRes
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		tmp := []int{}
		for i := 0; i < length; i++ {
			tree := queue[0]
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
			tmp = append(tmp, tree.Val)
			queue = queue[1:]
		}
		res = append(res, tmp)
	}
	return res
}

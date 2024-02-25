package binaryTree

// 判断二叉树是否是完全二叉树--广度遍历
// 若设二叉树的高度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第h层有叶子结点，并且叶子结点都是从左到右依次排布，这就是完全二叉树。
func isCompleteTree(tree *TreeNode) bool {
	if tree == nil {
		return false
	}
	queue := []*TreeNode{tree}
	end := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			end = true
		} else {
			if end {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}

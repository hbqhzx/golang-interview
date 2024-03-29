package binaryTree

// 还原二叉树
// 从前和中还原二叉树
// 对于任意一颗树而言，前序遍历的形式总是
//
// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// 即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
//
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
// 只要我们在中序遍历中定位到根节点，那么我们就可以分别知道左子树和右子树中的节点数目。由于同一颗子树的前序遍历和中序遍历的长度显然是相同的，因此我们就可以对应到前序遍历的结果中，对上述形式中的所有左右括号进行定位。
//
// 这样以来，我们就知道了左子树的前序遍历和中序遍历结果，以及右子树的前序遍历和中序遍历结果，我们就可以递归地对构造出左子树和右子树，再将这两颗子树接到根节点的左右位置。
func buildTreeByPreAndIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	midVal := preorder[0]
	root := &TreeNode{
		Val:   midVal,
		Left:  nil,
		Right: nil,
	}
	index := 0
	for i, v := range inorder {
		if v == midVal {
			index = i
		}
	}

	root.Left = buildTreeByPreAndIn(preorder[1:index+1], inorder[:index])
	root.Right = buildTreeByPreAndIn(preorder[index+1:], inorder[index+1:])
	return root

}

//从前和后还原二叉树
//前序和后序在本质上都是将父节点与子结点进行分离，但并没有指明左子树和右子树的能力，因此得到这两个序列只能明确父子关系，而不能确定一个二叉树。 故此法无。不能唯一确定一个二叉树。

// 从中和后还原二叉树
// 后序遍历的末尾元素即为当前子树的根节点
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	midVal := postorder[len(postorder)-1]
	root := &TreeNode{
		Val:   midVal,
		Left:  nil,
		Right: nil,
	}

	index := 0
	for i, v := range inorder {
		if v == midVal {
			index = i
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

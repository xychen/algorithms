package problem33

// 剑指 Offer 33. 二叉搜索树的后序遍历序列

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	splitIndex := -1
	// 找到左边的
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > root {
			break
		}
		splitIndex = i
	}
	// check右边的
	for i := splitIndex + 1; i < len(postorder)-1; i++ {
		if postorder[i] < root {
			return false
		}
	}
	// check子树
	left, right := true, true
	// 有左子树
	if splitIndex >= 0 {
		left = verifyPostorder(postorder[:splitIndex+1])
	}
	// 有右子树
	if splitIndex+1 < len(postorder)-1 {
		right = verifyPostorder(postorder[splitIndex+1 : len(postorder)-1])
	}
	return left && right
}

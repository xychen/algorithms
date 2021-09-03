package problem0106

// 从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	index := getIndex(inorder, rootVal)
	leftLen := index
	rightLen := len(inorder) - index - 1
	root.Left = buildTree(inorder[0:leftLen], postorder[0:leftLen])
	root.Right = buildTree(inorder[index+1:], postorder[leftLen:(leftLen+rightLen)])
	return root
}

func getIndex(inorder []int, value int) int {
	for i, v := range inorder {
		if v == value {
			return i
		}
	}
	return -1
}

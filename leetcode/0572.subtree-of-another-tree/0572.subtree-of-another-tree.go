package problem0572

// 另一棵树的子树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return checkSubtree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkSubtree(root *TreeNode, subRoot *TreeNode) bool {
	//同时为null
	if root == nil && subRoot == nil {
		return true
	}
	//只有一个为null
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val {
		return checkSubtree(root.Left, subRoot.Left) && checkSubtree(root.Right, subRoot.Right)
	}
	return false
}

package problem0965

//单值二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return traverse(root, root.Val)

}

func traverse(root *TreeNode, v int) bool {
	if root == nil {
		return true
	}
	if root.Val != v {
		return false
	}
	return traverse(root.Left, v) && traverse(root.Right, v)
}

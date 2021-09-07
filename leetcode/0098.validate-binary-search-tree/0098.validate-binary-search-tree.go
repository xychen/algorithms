package problem0098

// 验证二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	if !check(root.Left, min, root) {
		return false
	}
	if !check(root.Right, root, max) {
		return false
	}
	return true
}

package problem0814

// 二叉树剪枝

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	return doPrune(root)
}

func doPrune(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = doPrune(root.Left)
	root.Right = doPrune(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

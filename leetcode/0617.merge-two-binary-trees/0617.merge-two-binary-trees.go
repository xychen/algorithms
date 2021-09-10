package problem0617

// 合并二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	traversal(root1, root2)
	return root1
}

func traversal(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
	}
	if root1 == nil {
		root1 = &TreeNode{Val: root2.Val}
	}
	var r2l *TreeNode
	var r2r *TreeNode
	if root2 != nil {
		r2l, r2r = root2.Left, root2.Right
	}
	root1.Left = traversal(root1.Left, r2l)
	root1.Right = traversal(root1.Right, r2r)

	return root1
}

package problem0236

// 最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	n1 := lowestCommonAncestor(root.Left, p, q)
	n2 := lowestCommonAncestor(root.Right, p, q)
	if n1 != nil && n2 != nil {
		return root
	} else if n1 == nil {
		return n2
	}
	return n1
}

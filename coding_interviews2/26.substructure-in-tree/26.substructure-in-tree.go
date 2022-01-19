package problem26

// 剑指 Offer 26. 树的子结构

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return check(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func check(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return check(A.Left, B.Left) && check(A.Right, B.Right)
}

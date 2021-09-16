package problem0687

// 最长同值路径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxPath int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPath = 0
	singlePathLength(root)
	return maxPath
}

func singlePathLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLength := singlePathLength(root.Left)
	rightLength := singlePathLength(root.Right)
	leftLenWithRoot, rightLenWithRoot := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		leftLenWithRoot = leftLength + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rightLenWithRoot = rightLength + 1
	}
	if leftLenWithRoot+rightLenWithRoot > maxPath {
		maxPath = leftLenWithRoot + rightLenWithRoot
	}
	return max(leftLenWithRoot, rightLenWithRoot)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

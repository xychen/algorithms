package problem0124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT_MIN = -1 << 31

var ans int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxPathSum(root *TreeNode) int {
	ans = INT_MIN
	oneSideMax(root)
	return ans
}

func oneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, oneSideMax(root.Left))
	right := max(0, oneSideMax(root.Right))
	tmpVal := left + right + root.Val
	ans = max(ans, tmpVal)

	return max(left, right) + root.Val
}

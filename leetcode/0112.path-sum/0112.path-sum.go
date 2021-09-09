package problem0112

//路径和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return doPathSum(root, targetSum)
}

func doPathSum(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	if root.Left != nil && doPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if root.Right != nil && doPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

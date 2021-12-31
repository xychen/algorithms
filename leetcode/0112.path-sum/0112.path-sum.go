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
	res := false
	var backTrack func(root *TreeNode, sum int)
	backTrack = func(root *TreeNode, sum int) {
		if res || (root.Left == nil && root.Right == nil) {
			if sum == targetSum {
				res = true
			}
			return
		}
		if root.Left != nil {
			backTrack(root.Left, sum+root.Left.Val)
		}
		if root.Right != nil {
			backTrack(root.Right, sum+root.Right.Val)
		}
	}
	backTrack(root, root.Val)
	return res
}

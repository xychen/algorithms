package problem54

// 剑指 Offer 54. 二叉搜索树的第k大节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

package problem0144

//前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	traversal(root)
	return res
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	traversal(root.Left)
	traversal(root.Right)
}

package problem0145

//后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func postorderTraversal(root *TreeNode) []int {
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
	traversal(root.Left)
	traversal(root.Right)
	res = append(res, root.Val)
}

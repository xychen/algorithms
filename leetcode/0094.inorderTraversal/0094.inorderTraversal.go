package problem0094

//二叉树中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	doTraversal(root)
	return res
}

func doTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	doTraversal(root.Left)
	res = append(res, root.Val)
	doTraversal(root.Right)
}

package problem0404

//左叶子之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func sumOfLeftLeaves(root *TreeNode) int {
	res = 0
	traversal(root, false)
	return res
}

func traversal(root *TreeNode, isLeft bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && isLeft {
		res += root.Val
	}
	traversal(root.Left, true)
	traversal(root.Right, false)
}

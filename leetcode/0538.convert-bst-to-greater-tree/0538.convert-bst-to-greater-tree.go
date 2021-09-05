package problem0538

// 538. 把二叉搜索树转换为累加树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}

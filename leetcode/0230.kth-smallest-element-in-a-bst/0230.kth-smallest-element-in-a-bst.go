package problem0230

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var n int
var res int

func kthSmallest(root *TreeNode, k int) int {
	n, res = 0, -1
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	n++
	// fmt.Println(root.Val, ",n:", n)
	if k == n {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}

package problem0530

// 530. 二叉搜索树的最小绝对差
// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

// 差值是一个正数，其数值等于两值之差的绝对值。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	minDiff := int(^uint(0) >> 1)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		top := len(stack) - 1
		cur = stack[top]
		stack = stack[:top]

		if pre != nil {
			minDiff = min(minDiff, cur.Val-pre.Val)
		}
		pre = cur
		cur = cur.Right
	}
	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

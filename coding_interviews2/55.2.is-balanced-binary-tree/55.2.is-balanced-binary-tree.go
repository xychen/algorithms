package problem55_2

// 剑指 Offer 55 - II. 平衡二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var checkAndGetDepth func(node *TreeNode) (bool, int)
	checkAndGetDepth = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		f1, d1 := checkAndGetDepth(node.Left)
		if !f1 {
			return false, 0
		}
		f2, d2 := checkAndGetDepth(node.Right)
		if !f2 {
			return false, 0
		}
		if d1-d2 > 1 || d2-d1 > 1 {
			return false, 0
		}
		return true, max(d1, d2) + 1
	}
	f, _ := checkAndGetDepth(root)
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

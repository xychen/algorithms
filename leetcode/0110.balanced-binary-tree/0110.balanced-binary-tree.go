package problem0110

// 给定一个二叉树，判断它是否是高度平衡的二叉树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一
var ht map[*TreeNode]int

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if ht == nil {
		ht = make(map[*TreeNode]int)
	}

	if _, ok := ht[root.Left]; !ok {
		ht[root.Left] = high(root.Left)
	}
	if _, ok := ht[root.Right]; !ok {
		ht[root.Right] = high(root.Right)
	}
	if abs(ht[root.Left]-ht[root.Right]) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 方法二
func isBalanced2(root *TreeNode) bool {
	return high2(root) != -1
}

func high2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := high(root.Left)
	rh := high(root.Right)
	if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := high(root.Left)
	rh := high(root.Right)
	return max(lh, rh) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

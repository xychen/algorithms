package problem0104

// 104. 二叉树的最大深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层次遍历法
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		level++
		n := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return level
}

// 递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package problem0559

// N 叉树的最大深度

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	md := 0
	for _, n := range root.Children {
		md = max(maxDepth(n), md)
	}
	return md + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

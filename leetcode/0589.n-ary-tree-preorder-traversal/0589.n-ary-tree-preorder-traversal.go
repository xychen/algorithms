package problem0589

// 给定一个 N 叉树，返回其节点值的 前序遍历 。

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	traversal(root)
	return res
}

func traversal(root *Node) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	for _, Node := range root.Children {
		traversal(Node)
	}
}

package problem0590

//n叉树的后续遍历

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func postorder(root *Node) []int {
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
	for _, node := range root.Children {
		traversal(node)
	}
	res = append(res, root.Val)
}

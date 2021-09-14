package problem0623

// 在二叉树中增加一行

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return root
	}
	if depth == 1 {
		node := &TreeNode{Val: val}
		node.Left = root
		return node
	}
	queue := []*TreeNode{root}
	d := 1
	for (len(queue) > 0) && (d != depth-1) {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[l:]
		d++
	}

	for i := 0; i < len(queue); i++ {
		cur := queue[i]
		node1 := &TreeNode{Val: val}
		node1.Left = cur.Left
		cur.Left = node1
		node2 := &TreeNode{Val: val}
		node2.Right = cur.Right
		cur.Right = node2
	}
	return root
}

package problem32_2

// 剑指 Offer 32 - II. 从上到下打印二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		t := make([]int, 0)
		for i := 0; i < l; i++ {
			t = append(t, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, t)
		queue = queue[l:]
	}
	return res
}

package problem32_2

// 剑指 Offer 32 - III. 从上到下打印二叉树

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
	direct := 1
	for len(queue) > 0 {
		l := len(queue)
		t := make([]int, l)
		for i := 0; i < l; i++ {
			if direct > 0 {
				t[i] = queue[i].Val
			} else {
				t[l-i-1] = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}
		res = append(res, t)
		queue = queue[l:]
		direct = -1 * direct
	}
	return res
}

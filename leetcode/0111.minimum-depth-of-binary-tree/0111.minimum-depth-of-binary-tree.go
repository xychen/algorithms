package problem0111

//二叉树的最小深度
//解法一：二叉树的层次遍历（BFS）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//本质上是层次遍历
	depth := 1
	queue := make([]*TreeNode, 0)
	//第一个节点加入
	queue = append(queue, root)
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[i]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
		queue = queue[sz-1:]
	}
	return depth
}

package problem0951

// 翻转等价二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	queue1 := []*TreeNode{root1}
	queue2 := []*TreeNode{root2}
	for len(queue1) > 0 && len(queue2) > 0 {
		l := len(queue1)
		for i := 0; i < l; i++ {
			n1 := queue1[i]
			n2 := queue2[i]
			if !isEqual(n1, n2) {
				return false
			}
			switchChildren(n1, n2)
			//交换后，还是不相等
			if !isEqual(n1.Left, n2.Left) || !isEqual(n1.Right, n2.Right) {
				return false
			}
			if n1.Left != nil {
				queue1 = append(queue1, n1.Left)
				queue2 = append(queue2, n2.Left)
			}
			if n1.Right != nil {
				queue1 = append(queue1, n1.Right)
				queue2 = append(queue2, n2.Right)
			}
		}
		queue1 = queue1[l:]
		queue2 = queue2[l:]
	}
	if len(queue1) > 0 || len(queue2) > 0 {
		return false
	}
	return true
}

func switchChildren(root1 *TreeNode, root2 *TreeNode) {
	if !isEqual(root1.Left, root2.Left) {
		//交换
		tmp := root1.Right
		root1.Right = root1.Left
		root1.Left = tmp
	}
}

func isEqual(n1, n2 *TreeNode) bool {
	//都是nil
	if n1 == nil && n2 == nil {
		return true
	}
	//有一个是nil
	if n1 == nil || n2 == nil {
		return false
	}
	//都不是nil
	if n1.Val == n2.Val {
		return true
	}
	return false
}

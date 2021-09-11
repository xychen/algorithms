package problem0993

// 是否为堂兄弟节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || x == root.Val || y == root.Val {
		return false
	}
	queue := make([]*TreeNode, 0)
	//x,y的层级
	lx, ly := -1, -2
	//x,y的父节点
	var px *TreeNode
	var py *TreeNode
	xdone, ydone := false, false
	queue = append(queue, root)
	level := 1
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				if cur.Left.Val == x {
					px = cur
					lx = level + 1
					xdone = true
				}
				if cur.Left.Val == y {
					py = cur
					ly = level + 1
					ydone = true
				}
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				if cur.Right.Val == x {
					px = cur
					lx = level + 1
					xdone = true
				}
				if cur.Right.Val == y {
					py = cur
					ly = level + 1
					ydone = true
				}
			}
			if xdone && ydone {
				break
			}
		}
		queue = queue[qlen:]
		if xdone && ydone {
			break
		}
		level++
	}
	// fmt.Println(lx, ly, px.Val, py.Val)
	if !xdone || !ydone || lx != ly || px == py {
		return false
	}
	return true
}

package problem0701

// 二叉搜索树中的插入操作

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur, pre := root, root
	for cur != nil {
		pre = cur
		if cur.Val >= val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	node := &TreeNode{Val: val}
	if pre.Val >= val {
		pre.Left = node
	} else {
		pre.Right = node
	}
	return root
}

//递归法
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 方法二：递归法，好理解的，不够简洁
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	insert(root, val)
	return root
}

func insert(root *TreeNode, val int) {
	if root.Val == val {
		root.Val = val
		return
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return
		} else {
			insert(root.Left, val)
		}
	}

	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return
		} else {
			insert(root.Right, val)
		}
	}
}

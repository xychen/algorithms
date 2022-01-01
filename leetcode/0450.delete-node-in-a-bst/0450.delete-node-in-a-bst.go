package problem0450

// 删除二叉搜索树中的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 1.未找到
	// 2.找到了
	// 2.1 左右子节点都为空，直接删
	// 2.2 左子节点为空，右子节点补上
	// 2.3 右子节点为空，左子节点补上
	// 2.4 左右子节点都不为空，左子节点放到右子节点的最右侧，右子节点补位
	if root.Val == key {
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			p := root.Right
			for p.Left != nil {
				p = p.Left
			}
			p.Left = root.Left
			root = root.Right
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

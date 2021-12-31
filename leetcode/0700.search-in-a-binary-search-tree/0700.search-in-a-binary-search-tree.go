package problem0700

// 二叉搜索树中的搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

// 迭代法
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		}
	}
	return nil
}

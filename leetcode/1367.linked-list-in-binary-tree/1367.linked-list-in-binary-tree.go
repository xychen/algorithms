package problem1367

// 二叉树中的列表

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	//说明都验证完了
	if head == nil {
		return true
	}
	//说明head不为nil
	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	//说明都验证完了
	if head == nil {
		return true
	}
	//说明head不为nil
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, child := range children {
		//做选择
		res := dfs(head.Next, child)
		if res {
			return true
		}
		//回退
	}
	return false
}

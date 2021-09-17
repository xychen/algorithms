package problem1008

// 前序遍历构造二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	if len(preorder) == 0 {
		return root
	}
	i := getFirstGreaterIndex(root.Val, preorder)
	if i == 0 {
		root.Right = bstFromPreorder(preorder)
	} else if i < 0 {
		root.Left = bstFromPreorder(preorder)
	} else if i > 0 {
		root.Left = bstFromPreorder(preorder[0:i])
		root.Right = bstFromPreorder(preorder[i:])
	}
	return root
}

func getFirstGreaterIndex(v int, nums []int) int {
	res := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > v {
			res = i
			break
		}
	}
	return res
}

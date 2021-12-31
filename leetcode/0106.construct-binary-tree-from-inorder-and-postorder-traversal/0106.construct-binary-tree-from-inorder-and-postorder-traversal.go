package problem0106

// 从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}
	v := postorder[len(postorder)-1]
	index := getInorderIndex(inorder, v)
	root := &TreeNode{Val: v}
	root.Left = buildTree(inorder[0:index], postorder[0:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func getInorderIndex(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			return i
		}
	}
	return -1
}

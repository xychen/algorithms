package problem07

// 剑指 Offer 07. 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	i := findIndex(inorder, rootVal)
	root := &TreeNode{
		Val: rootVal,
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func findIndex(nums []int, val int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return -1
}

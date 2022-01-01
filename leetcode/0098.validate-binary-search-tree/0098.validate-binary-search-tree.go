package problem0098

// 验证二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：不用把方法二中的元素存到数组中，遍历过程中判断即可（保存前一个元素）
func isValidBST1_1(root *TreeNode) bool {
	var pre *TreeNode
	var checkBST func(root *TreeNode) bool
	checkBST = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := checkBST(root.Left)
		if pre != nil && pre.Val >= root.Val {
			return false
		}
		pre = root
		right := checkBST(root.Right)
		return left && right
	}
	return checkBST(root)
}

// 方法一（迭代法）：不用把方法二中的元素存到数组中，遍历过程中判断即可（保存前一个元素）
func isValidBST1_2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := len(stack) - 1
			cur = stack[top]
			stack = stack[:top]
			if pre != nil && pre.Val >= cur.Val {
				return false
			}
			pre = cur //保存前一个访问的节点
			cur = cur.Right
		}
	}
	return true
}

// 方法二：中序遍历，判断数组中是否递增
func isValidBST2(root *TreeNode) bool {
	inOrderNodes := make([]*TreeNode, 0)
	var inorderTraveral func(root *TreeNode)
	inorderTraveral = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraveral(root.Left)
		inOrderNodes = append(inOrderNodes, root)
		inorderTraveral(root.Right)
	}
	inorderTraveral(root)
	for i := 1; i < len(inOrderNodes); i++ {
		if inOrderNodes[i].Val <= inOrderNodes[i-1].Val {
			return false
		}
	}
	return true
}

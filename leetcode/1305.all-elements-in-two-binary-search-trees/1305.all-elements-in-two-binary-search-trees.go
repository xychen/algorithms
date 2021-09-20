package problem1305

// 两棵二叉搜索树中的所有元素
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历，然后归并
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	path1 := inorder(root1)
	path2 := inorder(root2)
	res := make([]int, len(path1)+len(path2))
	i, j, k := 0, 0, 0
	for i < len(path1) && j < len(path2) {
		if path1[i] < path2[j] {
			res[k] = path1[i]
			i++
		} else {
			res[k] = path2[j]
			j++
		}
		k++
	}
	for i < len(path1) {
		res[k] = path1[i]
		k++
		i++
	}
	for j < len(path2) {
		res[k] = path2[j]
		k++
		j++
	}
	return res
}

func inorder(root *TreeNode) []int {
	path := make([]int, 0)
	if root == nil {
		return path
	}
	pathLeft := inorder(root.Left)
	path = append(path, pathLeft...)
	path = append(path, root.Val)
	pathRight := inorder(root.Right)
	path = append(path, pathRight...)
	return path
}

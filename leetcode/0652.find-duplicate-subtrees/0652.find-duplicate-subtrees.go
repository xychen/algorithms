package problem0652

import (
	"strconv"
	"strings"
)

// 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
// 两棵树重复是指它们具有相同的结构以及相同的结点值。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []*TreeNode
var hm map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = make([]*TreeNode, 0)
	hm = make(map[string]int)
	traverse(root)
	return res
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	leftStr := traverse(root.Left)
	rightStr := traverse(root.Right)
	str := strings.Join([]string{leftStr, rightStr, strconv.Itoa(root.Val)}, ",")
	if num, ok := hm[str]; ok && num == 1 {
		res = append(res, root)
	}
	hm[str]++
	return str
}

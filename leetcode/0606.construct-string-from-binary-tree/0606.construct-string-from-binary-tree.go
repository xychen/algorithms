package problem0606

// 根据二叉树创建字符串
import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法二：
func tree2str2(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return res
	}
	if root.Right == nil {
		return res + "(" + tree2str2(root.Left) + ")"
	}
	return res + "(" + tree2str2(root.Left) + ")(" + tree2str2(root.Right) + ")"
}

//方法一：
func tree2str(root *TreeNode) string {
	if root == nil {
		return "()"
	}
	res := strconv.Itoa(root.Val)
	res += "(" + trimTailBracket(tree2str(root.Left)) + ")"
	res += "(" + trimTailBracket(tree2str(root.Right)) + ")"
	return trimTailBracket(res)
}

func trimTailBracket(str string) string {
	if str == "()" {
		return ""
	}
	if len(str) < 2 {
		return str
	}
	for len(str) >= 2 && str[len(str)-2:] == "()" {
		str = str[0 : len(str)-2]
	}
	return str
}

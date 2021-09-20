package problem1110

// 删点成林
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode
	if root == nil {
		return res
	}
	res = make([]*TreeNode, 0)
	if len(to_delete) <= 0 {
		res = append(res, root)
		return res
	}
	toDelMap := sliceToMap(to_delete)
	resMap := make(map[*TreeNode]bool)

	queue := []*TreeNode{root}
	resMap[root] = true
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			//当前节点要被删除
			if _, ok := toDelMap[cur.Val]; ok {
				delete(resMap, cur)
				if cur.Left != nil {
					resMap[cur.Left] = true
				}
				if cur.Right != nil {
					resMap[cur.Right] = true
				}
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				if _, ok := toDelMap[cur.Left.Val]; ok {
					//切断联系
					cur.Left = nil
				}
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				if _, ok := toDelMap[cur.Right.Val]; ok {
					//切断联系
					cur.Right = nil
				}
			}

		}
		queue = queue[l:]
	}
	return getKeys(resMap)
}

func sliceToMap(nums []int) map[int]bool {
	ht := make(map[int]bool)
	for _, num := range nums {
		ht[num] = true
	}
	return ht
}

func getKeys(ht map[*TreeNode]bool) []*TreeNode {
	res := make([]*TreeNode, len(ht))
	i := 0
	for k, _ := range ht {
		res[i] = k
		i++
	}
	return res
}

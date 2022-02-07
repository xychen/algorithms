package problem0763

// 763. 划分字母区间
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

func partitionLabels(s string) []int {
	// 字母最远的位置
	ht := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ht[s[i]] = i
	}
	res := make([]int, 0)
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, ht[s[i]])
		// 新的一组
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

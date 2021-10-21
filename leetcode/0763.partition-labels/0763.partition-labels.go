package problem0763

// 763. 划分字母区间
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

func partitionLabels(s string) []int {
	// 字母 => 最后位置的index
	mem := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mem[s[i]] = i
	}
	ans := make([]int, 0)
	cover := 0
	prelen := -1
	for i := 0; i < len(s); i++ {
		alph := s[i]
		cover = max(cover, mem[alph])
		if i == cover {
			ans = append(ans, i-prelen)
			cover = 0
			prelen = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

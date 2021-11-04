package problem05

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	c := 0
	// 统计空格数量
	for _, ch := range s {
		if ch == ' ' {
			c++
		}
	}
	res := make([]rune, len(s)+2*c)
	j := 0
	for _, ch := range s {
		if ch == ' ' {
			res[j] = '%'
			j++
			res[j] = '2'
			j++
			res[j] = '0'
		} else {
			res[j] = ch
		}
		j++
	}
	return string(res)
}

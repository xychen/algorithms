package problem1047

// 1047. 删除字符串中的所有相邻重复项

func removeDuplicates(s string) string {
	slist := []byte(s)
	stack := make([]byte, 0)

	for _, c := range slist {
		if len(stack) == 0 {
			// push
			stack = append(stack, c)
			continue
		}
		topV := stack[len(stack)-1]
		if c == topV {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			// push
			stack = append(stack, c)
		}
	}
	return string(stack)
}

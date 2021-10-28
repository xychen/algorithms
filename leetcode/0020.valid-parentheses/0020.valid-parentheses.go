package problem0020

// 20. 有效的括号

func isValid(s string) bool {
	stack := make([]byte, 0)
	slist := []byte(s)
	ht := map[byte]byte{
		byte('('): byte(')'),
		byte('{'): byte('}'),
		byte('['): byte(']'),
	}
	for _, c := range slist {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		top := stack[len(stack)-1]
		if v, ok := ht[top]; ok && v == c {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

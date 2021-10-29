package problem0150

// 150. 逆波兰表达式求值

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
			continue
		}
		a := stack[len(stack)-2]
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, a/b)
		}
	}
	return stack[len(stack)-1]
}

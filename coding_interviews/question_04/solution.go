package main

import "fmt"

//将字符串中的空格替换成%20
func ReplaceBlank(s string) string {
	n := len(s)
	blankCount := 0
	for i := 0; i < n; i++ {
		c := s[i]
		if c == 32 {
			blankCount += 1
		}
	}
	newLen := n + 2*blankCount
	strByte := make([]byte, newLen)
	op := n - 1
	newp := newLen - 1
	for newp >= 0 {
		if s[op] == 32 {
			strByte[newp] = '0'
			newp -= 1
			strByte[newp] = '2'
			newp -= 1
			strByte[newp] = '%'
		} else {
			strByte[newp] = s[op]
		}
		newp -= 1
		op -= 1
	}
	return string(strByte)
}

func main() {
	str := "we are happy"
	str = ReplaceBlank(str)
	fmt.Println(str)
}

package main

import "fmt"

func count(n uint8) int {
	var c uint8 = 0
	for n != 0 {
		c += n & 0x1
		n >>= 1
	}
	return int(c)
}

func main() {
	var a uint8 = 7
	fmt.Println(count(a))
}

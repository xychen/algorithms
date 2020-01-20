package main

import "fmt"

func find(ids []int) int {
	count := 0
	var condition int
	for _, id := range ids {
		if count == 0 {
			condition = id
			count = 1
		} else {
			if condition == id {
				count += 1
			} else {
				count -= 1
			}
		}
	}
	return condition
}

func main() {
	arr := []int{1, 2, 2, 2, 3, 3, 2}
	res := find(arr)
	fmt.Println(res)
}

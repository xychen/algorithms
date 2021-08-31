package problem0752

import (
	"strconv"
	"strings"
)

var deadendsMap map[string]bool
var visit map[string]bool

func openLock(deadends []string, target string) int {
	deadendsMap = make(map[string]bool)
	for _, item := range deadends {
		deadendsMap[item] = true
	}
	visit = make(map[string]bool)
	return BFS(target)
}

func plusOne(str string, j int) string {
	s := strings.Split(str, "")
	num, _ := strconv.Atoi(s[j])
	num = (num + 1) % 10
	s[j] = strconv.Itoa(num)
	return strings.Join(s, "")
}

func minusOne(str string, j int) string {
	s := strings.Split(str, "")
	num, _ := strconv.Atoi(s[j])
	num = (num - 1 + 10) % 10
	s[j] = strconv.Itoa(num)
	return strings.Join(s, "")
}

func BFS(target string) int {
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visit["0000"] = true
	step := 0
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[i]
			if _, ok := deadendsMap[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				pone := plusOne(cur, j)
				if _, ok := visit[pone]; !ok {
					queue = append(queue, pone)
					visit[pone] = true
				}
				mone := minusOne(cur, j)
				if _, ok := visit[mone]; !ok {
					queue = append(queue, mone)
					visit[mone] = true
				}
			}
		}
		//本次循环遍历过的元素出队列
		queue = queue[sz:]
		step++
	}
	return -1
}

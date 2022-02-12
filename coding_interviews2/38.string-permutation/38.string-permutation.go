package problem38

import (
	"fmt"
	"sort"
)

// 方法一：下一个组合
func permutation(s string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	strbyte := ([]byte)(s)
	sort.Slice(strbyte, func(i, j int) bool {
		if strbyte[i] < strbyte[j] {
			return true
		}
		return false
	})
	res := []string{string(strbyte)}
	n := len(s)
	var nextzuhe func() bool
	nextzuhe = func() bool {
		// 从后往前找到第一个较小的元素
		i := n - 2
		for i >= 0 && strbyte[i] >= strbyte[i+1] {
			i--
		}
		// 没有下个组合
		if i < 0 {
			return false
		}
		j := n - 1
		for j > i && strbyte[j] <= strbyte[i] {
			j--
		}
		strbyte[i], strbyte[j] = strbyte[j], strbyte[i]
		left, right := i+1, n-1
		for left < right {
			strbyte[left], strbyte[right] = strbyte[right], strbyte[left]
			left++
			right--
		}
		res = append(res, string(strbyte))
		return true
	}
	for {
		if !nextzuhe() {
			break
		}
	}
	return res
}

// 方法二：排序+回溯
func permutation2(s string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	res := make([]string, 0)
	strbyte := ([]byte)(s)
	// 排序
	sort.Slice(strbyte, func(i, j int) bool {
		return strbyte[i] < strbyte[j]
	})
	fmt.Println(string(strbyte))
	vis := make([]bool, len(s))
	var backtrate func(path []byte)
	backtrate = func(path []byte) {
		if len(path) == len(strbyte) {
			res = append(res, string(path))
			return
		}
		for i, b := range strbyte {
			if vis[i] || (i > 0 && strbyte[i] == strbyte[i-1] && !vis[i-1]) {
				continue
			}
			// 做选择
			vis[i] = true
			path = append(path, b)
			backtrate(path)
			// 回退
			path = path[:len(path)-1]
			vis[i] = false
		}
	}
	path := make([]byte, 0)
	backtrate(path)
	return res
}

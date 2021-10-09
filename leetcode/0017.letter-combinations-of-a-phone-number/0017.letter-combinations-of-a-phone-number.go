package problem0017

// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

import "strings"

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	n := len(digits)
	if n == 0 {
		return res
	}
	input := strings.Split(digits, "")
	strmap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var solve func(starti int, path []string)
	solve = func(starti int, path []string) {
		if len(path) == n {
			tmp := strings.Join(path, "")
			res = append(res, tmp)
			return
		}
		for i := starti; i < n; i++ {
			k := input[i]
			for j := 0; j < len(strmap[k]); j++ {
				path = append(path, strmap[k][j])
				solve(i+1, path)
				path = path[:len(path)-1]
			}

		}
	}
	path := make([]string, 0)
	solve(0, path)
	return res
}

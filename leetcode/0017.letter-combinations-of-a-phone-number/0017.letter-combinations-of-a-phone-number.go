package problem0017

// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dmap := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	var backtrack func(digits string, starIndex int, path []byte)
	backtrack = func(digits string, starIndex int, path []byte) {
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}
		numChar := digits[starIndex]
		for _, char := range dmap[numChar] {
			// 做选择
			path = append(path, char)
			backtrack(digits, starIndex+1, path)
			// 回退
			path = path[:len(path)-1]
		}
	}
	path := make([]byte, 0)
	backtrack(digits, 0, path)
	return res
}

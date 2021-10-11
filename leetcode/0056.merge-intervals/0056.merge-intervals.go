package problem0056

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

import "sort"

func merge(intervals [][]int) [][]int {
	// 按start升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	curStart, curEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 有交集
		if intervals[i][0] <= curEnd {
			curEnd = max(intervals[i][1], curEnd)
		} else {
			res = append(res, []int{curStart, curEnd})
			curStart, curEnd = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{curStart, curEnd})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

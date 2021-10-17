package problem0435

import "sort"

// 435. 无重叠区间
// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
// 注意:
// 可以认为区间的终点总是大于它的起点。
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

// 贪心算法
func eraseOverlapIntervals(intervals [][]int) int {
	// 按右边界升序排
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 思路：选择没有交集的留下，总集合数-留下的数 => 移除的数量

	sel := 1
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// 有交集(起始位置<当前的结束位置)
		if intervals[i][0] < cur[1] {
			continue
		}
		sel++
		cur = intervals[i]
	}
	return len(intervals) - sel
}

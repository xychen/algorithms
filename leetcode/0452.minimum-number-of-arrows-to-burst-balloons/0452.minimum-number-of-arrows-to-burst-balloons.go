package problem0452

import "sort"

// 452. 用最少数量的箭引爆气球
// 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。

// 一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。

// 给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。

func findMinArrowShots(points [][]int) int {
	// 按照xend 递增排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// 有n个不重叠的区间，就需要放n键
	sel := 1
	cur := points[0]
	for i := 1; i < len(points); i++ {
		// 有交集，不能选 （边界触碰也算有交集）
		if points[i][0] <= cur[1] {
			continue
		}
		sel++
		cur = points[i]
	}
	return sel
}

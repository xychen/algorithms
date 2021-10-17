package problem0406

// 406. 根据身高重建队列
// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。

// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		}
		if people[i][0] == people[j][0] {
			if people[i][1] < people[j][1] {
				return true
			} else {
				return false
			}
		}
		return false
	})

	ans := make([][]int, len(people))
	for _, p := range people {
		position := p[1]
		copy(ans[position+1:], ans[position:])
		ans[position] = p
	}
	return ans
}

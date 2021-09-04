package problem0398

// 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

import "math/rand"

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{Nums: nums}
}

func (this *Solution) Pick(target int) int {
	i := 0
	res := 0
	for j, n := range this.Nums {
		if n == target {
			i++
			if rand.Intn(i) == 0 {
				res = j
			}
		}
	}
	return res
}

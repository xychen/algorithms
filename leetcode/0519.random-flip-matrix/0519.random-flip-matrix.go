package problem0519

// 题中给出一个 n_rows 行 n_cols 列的二维矩阵，且所有值被初始化为 0。要求编写一个 flip 函数，均匀随机的将矩阵中的 0 变为 1，并返回该值的位置下标 [row_id,col_id]；
//  同样编写一个 reset 函数，将所有的值都重新置为 0。尽量最少调用随机函数 Math.random()，并且优化时间和空间复杂度。

import "math/rand"

type Solution struct {
	row    int
	column int
	ht     map[int]int
	remain int
}

func Constructor(m int, n int) Solution {
	return Solution{
		row:    m,
		column: n,
		ht:     make(map[int]int),
		remain: m * n,
	}
}

func (this *Solution) Flip() []int {
	randomV := rand.Intn(this.remain)
	if _, ok := this.ht[randomV]; !ok {
		this.ht[randomV] = randomV
	}
	realIndex := this.ht[randomV]
	this.remain--
	//交换位置(左侧是值为0的， 右侧是值为1)
	if _, ok := this.ht[this.remain]; !ok {
		this.ht[this.remain] = this.remain
	}
	this.ht[randomV] = this.ht[this.remain]

	return []int{realIndex / this.column, realIndex % this.column}
}

func (this *Solution) Reset() {
	this.ht = make(map[int]int)
	this.remain = this.row * this.column
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */

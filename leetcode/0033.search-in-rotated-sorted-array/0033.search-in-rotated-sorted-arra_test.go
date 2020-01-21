package problem0033

import "testing"

type param struct {
	nums   []int
	target int
}

type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

var qs = []question{
	question{p: param{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, a: ans{res: 4}},
	question{p: param{nums: []int{5, 1, 3}, target: 3}, a: ans{res: 2}},
	question{p: param{nums: []int{3, 1}, target: 1}, a: ans{res: 1}},
	question{p: param{nums: []int{3, 1}, target: 3}, a: ans{res: 0}},
}

func Test_Search1(t *testing.T) {
	for _, q := range qs {
		if q.a.res != Search1(q.p.nums, q.p.target) {
			t.Errorf("Search1输入:%v", q.p)
		}
	}
}

func Test_Search2(t *testing.T) {
	for _, q := range qs {
		if q.a.res != Search2(q.p.nums, q.p.target) {
			t.Errorf("Search2输入:%v", q.p)
		}
	}
}

func Test_Search3(t *testing.T) {
	for _, q := range qs {
		if q.a.res != Search3(q.p.nums, q.p.target) {
			t.Errorf("Search3输入:%v", q.p)
		}
	}
}

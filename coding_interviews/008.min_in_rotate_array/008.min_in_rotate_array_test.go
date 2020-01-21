package problem008

import "testing"

type param struct {
	one []int
}

type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

//旋转数组中的最小值
func Test_MinInRotateArr(t *testing.T) {
	qs := []question{
		question{p: param{one: []int{3, 4, 5, 1, 2}}, a: ans{res: 1}},
		question{p: param{one: []int{3, 4, 5, 6, 1, 2}}, a: ans{res: 1}},
		question{p: param{one: []int{1, 2, 3, 4, 5, 6}}, a: ans{res: 1}},
	}

	for _, q := range qs {
		if q.a.res != MinInRotateArr(q.p.one) {
			t.Errorf("输入：%v", q.p.one)
		}
	}
}

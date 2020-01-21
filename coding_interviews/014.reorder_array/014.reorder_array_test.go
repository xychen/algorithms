package problem014

import "testing"

type param struct {
	nums []int
}

type question struct {
	p param
}

//检测前半部分是奇数，后半部分是偶数
func check(nums []int) bool {
	i := 0
	for i < len(nums) && nums[i]&0x1 == 1 {
		i += 1
	}
	for i < len(nums) && nums[i]&0x1 == 0 {
		i += 1
	}
	if i != len(nums) {
		return false
	}
	return true
}

func Test_ReorderArray(t *testing.T) {
	qs := []question{
		question{p: param{nums: []int{1, 2, 3, 4, 5, 6}}},
		question{p: param{nums: []int{1, 1, 1, 1, 1, 2}}},
		question{p: param{nums: []int{1, 1, 1, 1, 1, 1}}},
		question{p: param{nums: []int{1, 2, 2, 2, 2, 2}}},
		question{p: param{nums: []int{2, 2, 2, 2, 2, 2}}},
	}

	for _, q := range qs {
		if !check(ReorderArray(q.p.nums)) {
			t.Errorf("输入:%v", q.p.nums)
		}
	}
}

package problem007

import (
	"testing"
)

func Test_OK(t *testing.T) {
	cq1 := NewCQueue()
	_, err := cq1.DeleteHead()
	if err == nil {
		t.Errorf("删除空队列错误")
	}
	cq1.AppendTail(1)
	cq1.AppendTail(2)
	cq1.AppendTail(3)
	tmp, err := cq1.DeleteHead()
	if err != nil || tmp != 1 {
		t.Errorf("队列操作错误1")
	}

	cq1.AppendTail(4)
	tmp, err = cq1.DeleteHead()
	if err != nil || tmp != 2 {
		t.Errorf("队列操作错误2")
	}

	tmp, err = cq1.DeleteHead()
	if err != nil || tmp != 3 {
		t.Errorf("队列操作错误3")
	}

	tmp, err = cq1.DeleteHead()
	if err != nil || tmp != 4 {
		t.Errorf("队列操作错误4")
	}
}

package problem0752

import (
	"testing"
)

func Test_plusOne(t *testing.T) {
	res := plusOne("0000", 1)
	if res != "0100" {
		t.Errorf("plusOne error")
	}
}

func Test_minusOne(t *testing.T) {
	res := minusOne("0000", 1)
	if res != "0900" {
		t.Errorf("minusOne error")
	}
}

func Test_BFS(t *testing.T) {
	deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target := "8888"
	res := openLock(deadends, target)
	if res != -1 {
		t.Errorf("openLock error")
	}
}

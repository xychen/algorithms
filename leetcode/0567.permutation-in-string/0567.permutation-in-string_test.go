package problem0567

import (
	"testing"
)

type input struct {
	s1  string
	s2  string
	res bool
}

var testCase = []input{
	{"intention", "execution", false},
	{"trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine", true},
	{"hello", "ooolleoooleh", false},
	{"ab", "eidboaoo", false},
}

func Test_checkInclusion(t *testing.T) {
	for _, item := range testCase {
		if item.res != checkInclusion(item.s1, item.s2) {
			t.Errorf("checkInclusion error - %s:%s", item.s1, item.s2)
		}
		if item.res != checkInclusion2(item.s1, item.s2) {
			t.Errorf("checkInclusion2 error - %s:%s", item.s1, item.s2)
		}
	}
}

func Test_checkInclusion2(t *testing.T) {

	for _, item := range testCase {
		if item.res != checkInclusion2(item.s1, item.s2) {
			t.Errorf("checkInclusion2 error - %s:%s", item.s1, item.s2)
		}
	}
}

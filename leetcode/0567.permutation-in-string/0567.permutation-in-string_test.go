package problem0567

import (
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	type input struct {
		s1  string
		s2  string
		res bool
	}
	testCase := []input{
		{"intention", "execution", false},
		{"trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine", true},
		{"hello", "ooolleoooleh", false},
	}
	for _, item := range testCase {
		if item.res != checkInclusion(item.s1, item.s2) {
			t.Errorf("checkInclusion error - %s:%s", item.s1, item.s2)
		}
		// res := checkInclusion(item.s1, item.s2)
		// fmt.Println(res)
	}
}

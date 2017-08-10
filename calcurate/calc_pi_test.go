package calcurate

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		desc string
		cmd  string
	}{
		{"未作成の計算", "noCmd"},
		{"円周率を計算", "pi"},
	}
	for _, test := range testCases {
		fmt.Println(Calc(test.cmd))
	}
}

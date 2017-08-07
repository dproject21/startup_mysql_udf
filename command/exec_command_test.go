package command

import (
	"testing"
)

func TestExec(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		want string
	}{
		{"OSコマンドの実行", "hostname", "mbp2015.local\n"},
	}
	for _, test := range testCases {
		out, err := Exec(test.in)
		if err != nil {
			t.Errorf("ExecError : TestCase %s, Message %s", test.desc, err)
		}
		if out != test.want {
			t.Errorf("ResultError : TestCase %s, Result %s", test.desc, out)
		}
	}
}

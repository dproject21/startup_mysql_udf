package command

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	testCases := []struct {
		desc string
		in   []string
		want string
	}{
		{"OSコマンドの実行", []string{"hostname"}, "mbp2015.local\n"},
		{"OSコマンドの実行", []string{"ls", "-la"}, "test"},
	}
	for _, test := range testCases {
		out, err := Exec(test.in)
		if err != nil {
			t.Errorf("ExecError : TestCase %s, Message %s", test.desc, err)
		}
		fmt.Println(out)
	}
}

package command

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
	}{
		{"OSコマンドの実行", "hostname"},
		{"OSコマンドの実行", "ls -la"},
		{"OSコマンドの実行", "touch test.txt"},
		{"OSコマンドの実行", "ls -la"},
		{"OSコマンドの実行", "rm -f test.txt"},
		{"OSコマンドの実行", "ls -la"},
	}
	for _, test := range testCases {
		out, err := Exec(test.in)
		if err != nil {
			t.Errorf("ExecError : TestCase %s, Message %s", test.desc, err)
		}
		fmt.Println(out)
	}
}

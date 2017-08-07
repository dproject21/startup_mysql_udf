package command

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func Exec(args string) (string, error) {
	c, err := shellwords.Parse(args)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var out []byte
	switch len(args) {
	case 0:
		out = make([]byte, 0, 0)
	case 1:
		out, err = exec.Command(c[0]).Output()
	default:
		out, err = exec.Command(c[0], c[1:]...).Output()
	}
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

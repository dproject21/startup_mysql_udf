package command

import (
	"fmt"
	"os/exec"
)

func Exec(args []string) (string, error) {
	var out []byte
	var err error
	switch len(args) {
	case 0:
		out = make([]byte, 0, 0)
	case 1:
		out, err = exec.Command(args[0]).Output()
	default:
		out, err = exec.Command(args[0], args[1:]...).Output()
	}
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

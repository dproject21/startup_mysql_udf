package command

import (
	"fmt"
	"os/exec"
)

func Exec(args string) (string, error) {
	out, err := exec.Command(args).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

package command

import (
	"fmt"
	"os/exec"
)

func Exec(cmd string) (string, error) {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

package utils

import (
	"os/exec"
	"strings"
)

func Execcmd(cmdStr string) (ret string, erro error) {
	args := strings.Split(cmdStr, " ")
	cmd := exec.Command(args[0], args[1:]...)
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(res), err

}

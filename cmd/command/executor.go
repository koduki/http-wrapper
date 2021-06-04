package command

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(command string, args []string) string {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	var status string
	if err == nil {
		status = "success"
	} else {
		status = "error"
		fmt.Println(err)
	}

	return status
}

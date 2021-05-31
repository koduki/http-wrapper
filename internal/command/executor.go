package command

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec() string {
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	var msg string
	if err == nil {
		msg = "success"
	} else {
		msg = "error"
		fmt.Println(err)
	}

	return msg
}

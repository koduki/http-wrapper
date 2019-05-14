package main

import "fmt"
import "os"
import "os/exec"
import "net/http"

func main() {
	port := os.Args[1]
	cmd := os.Args[2]
	args := os.Args[3:]
	http.HandleFunc("/", handler(cmd, args))
	http.ListenAndServe(":" + port, nil)
}

func handler(cmd string, args []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := run(cmd, args)
		fmt.Fprintf(w, msg)
	}
}

func run(command string, args []string) string {
	
	cmd := exec.Command(command, args...)
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
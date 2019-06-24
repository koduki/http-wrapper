package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

type Options struct {
	port int
}

var (
	o = &Options{}
)

var RootCmd = &cobra.Command{
	Use:   "hwrap [flags] command",
	Short: "HTTP server wrapper for command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		xs := scan(args[0])
		command := xs[0]
		cmdArgs := xs[1:len(xs)]

		fmt.Printf("port:%d, commad:%s, args:%s\n", o.port, command, cmdArgs)

		http.HandleFunc("/", handler(command, cmdArgs))
		http.ListenAndServe(":"+strconv.Itoa(o.port), nil)
	},
}

func init() {
	RootCmd.Flags().IntVarP(&o.port, "port", "p", 8080, "port number")
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

func scan(str string) []string {
	var head []string
	var tail string

	if str == "" {
		return []string{}
	} else {
		x := string(str[0])
		switch x {
		case "'":
			head, tail = splitQuote(str)
		case "\"":
			head, tail = splitDQuote(str)
		case " ":
			head, tail = skipSpace(str)
		default:
			head, tail = splitSpace(str)
		}
		return append(head, scan(tail)...)
	}
}

func skipSpace(str string) ([]string, string) {
	return []string{}, str[1:len(str)]
}

func splitSpace(str string) ([]string, string) {
	head := ""
	tail := ""
	for i := 0; i < len(str); i++ {
		x := string(str[i])
		if x == " " {
			tail = str[(i + 1):len(str)]
			break
		} else {
			head += x
		}
	}
	return []string{head}, tail
}

func splitQuote(str string) ([]string, string) {
	head := ""
	tail := ""
	for i := 1; i < len(str); i++ {
		x := string(str[i])
		if x == "'" {
			tail = str[i+1 : len(str)]
			break
		} else {
			head += x
		}
	}
	return []string{head}, tail
}

func splitDQuote(str string) ([]string, string) {
	head := ""
	tail := ""
	for i := 1; i < len(str); i++ {
		x := string(str[i])
		if x == "\"" {
			tail = str[i+1 : len(str)]
			break
		} else {
			head += x
		}
	}
	return []string{head}, tail
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// os.Input
// os.Output
// better: https://pkg.go.dev/github.com/nsf/termbox-go#Event

func main() {
	sh := NewShell(os.Stdin, os.Stdout)
	sh.Greeteings()
	for {
		sh.Listen()
	}
}

func NewShell(r io.Reader, w io.Writer) *GoShell {
	return &GoShell{r, w}
}

// 1 Greeteings
// 2 pwd
// 3 command line $

type GoShell struct {
	r io.Reader
	w io.Writer
}

func (sh *GoShell) Listen() {
	sh.line()

	// read from r
	buf := bufio.NewScanner(sh.r)
	if !buf.Scan() {
		os.Exit(0)
	}
	str := buf.Text()

	commands := strings.Split(str, " ")

	switch commands[0] {
	case "pwd":
		writeString(sh.w, pwd(), "\n")

	case "cd":
		if len(commands) == 2 {
			cd(commands[1])
		} else {
			cd("")
		}

	case "ls":
		writeString(sh.w, ls(pwd()))

	case "nc":
		nc(sh.r)

	case "echo":
		echo(sh.w, commands)

	case "exit":
		os.Exit(0)

	default:
		writeString(sh.w, fmt.Sprintf("Unknown command [%s]\n", str))
	}
}

func (sh *GoShell) Greeteings() {
	writeString(sh.w, fmt.Sprintf("Welcome %s!\n", username()))
}

// print command line
func (sh *GoShell) line() {
	writeString(sh.w, fmt.Sprintf("[%s:%s]$ ", username(), pwd()))
}

// ############################################################### //

func cd(path string) {
	if path == "" || path == "~" {
		path = os.Getenv("HOME")
	}

	if err := os.Chdir(path); err != nil {
		log.Println(err)
	}
}

func ls(path string) string {
	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return ""
	}

	str := ""

	for _, v := range dirs {
		str += v.Name() + "\n"
	}

	return str
}

func username() string {
	return os.Getenv("USER")
}

func echo(w io.Writer, strs []string) {

	for i, v := range strs {
		if len(v) > 2 && v[0] == '$' {
			strs[i] = os.Getenv(v[1:])
		}
	}

	strs = append(strs, "\n")
	writeString(w, strs...)
}

// return current path
func pwd() string {
	if pwd, err := os.Getwd(); err != nil {
		log.Println(err)
		return ""
	} else {
		return pwd
	}
}

// write string to writer
func writeString(w io.Writer, str ...string) {
	io.WriteString(w, strings.Join(str, " "))
}

// -u udp, default tcp
// hostname port
func nc(r io.Reader) {

	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Println(err)
		return
	}

	io.WriteString(conn, "bla bla")
}

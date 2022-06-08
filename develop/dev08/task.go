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

// слушать команды
//

// os.Input
// os.Output

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
		sh.pwd()
	case "cd":
		if len(commands) == 2 {
			sh.cd(commands[1])
		}
	case "nc":
		nc(sh.r)

	case "exit":
		os.Exit(0)

	default:
		writeString(sh.w, fmt.Sprintf("Unknown command [%s]\n", str))
	}
}

// print current path
func (sh *GoShell) pwd() {
	io.WriteString(sh.w, pwd()+"\n")
}

func (sh *GoShell) cd(path string) {
	switch path {
	case "":
		log.Println("to home dir")
	default:
		log.Println("to path", path)
	}

	// .. up
	// ./ relative
	// /../../../ absolute
	// . current
}

func (sh *GoShell) Greeteings() {
	writeString(sh.w, fmt.Sprintf("Welcome %s!\n", username()))
}

// print command line
func (sh *GoShell) line() {
	writeString(sh.w, fmt.Sprintf("[%s:%s]$ ", username(), pwd()))
}

// ############################################################### //

func username() string {
	return os.Getenv("USER")
}

// return current path
func pwd() string {
	return os.Getenv("PWD")
}

// write string to writer
func writeString(w io.Writer, str string) {
	io.WriteString(w, str)
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

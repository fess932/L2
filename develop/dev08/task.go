package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
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

func (sh *GoShell) Greeteings() {
	io.WriteString(sh.w, "Hello, %username%")
}

// print command line
func (sh *GoShell) line() {
	io.WriteString(sh.w, "\n[%username%]$ ")
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
	buf.Scan()
	str := buf.Text()

	switch str {
	case "pwd":
		sh.pwd()
	case "nc":
		nc(sh.r)

	case "exit":
		os.Exit(0)
	}

	io.WriteString(sh.w, str)

}

// print current path
func (sh *GoShell) pwd() {
	io.WriteString(sh.w, os.Getenv("PWD"))
}

func (sh *GoShell) cd() {
	// .. up
	// ./ relative
	// /../../../ absolute
	// . current
}

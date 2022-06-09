package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
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
	greeteings(sh.w)

	for {
		sh.Input()
	}
}

func NewShell(r io.Reader, w io.Writer) *GoShell {
	return &GoShell{
		r, w,
		map[string]InitCommand{
			"cd":   CD,
			"pwd":  PWD,
			"ls":   LS,
			"echo": Echo,
			"ps":   PS,
			"kill": Kill,
			"exec": Exec,
		},
	}
}

type GoShell struct {
	r io.Reader
	w io.Writer

	commandTable map[string]InitCommand
}

type Command struct {
	Name string
	Args []string

	r io.Reader
	w io.Writer
}

func (sh *GoShell) parseCommands(str string, r io.Reader, w io.Writer) (cmds []func()) {
	cs := strings.Split(str, "|")

	// pipe magic
	// cmd1 write to out1, out1 as in1 to cmd2, cmd2 write to stdout

	var tmp []string
	for i, v := range cs {
		tmp = strings.Split(v, " ")
		cmd := Command{Name: tmp[0], Args: tmp[1:]}

		if i == 0 {
			cmd.w = w
		}

		if i == len(cs)-1 {
			cmd.r = r
		}

		if cmdr, ok := sh.commandTable[cmd.Name]; ok {
			cmds = append(cmds, cmdr(cmd))

			continue
		}

		cmds = append(cmds, Fork(cmd))
	}

	return cmds
}

func (sh *GoShell) Input() {
	line(sh.w)

	// read from r
	buf := bufio.NewScanner(sh.r)
	if !buf.Scan() {
		os.Exit(0)
	}

	str := buf.Text()
	if len(str) == 0 {
		return
	}

	cmds := sh.parseCommands(str, sh.r, sh.w)

	if len(cmds) > 1 {
		log.Println("OOPS, TODO implement pipes")

		return
	}

	cmds[0]()
}

func greeteings(w io.Writer) {
	writeString(w, fmt.Sprintf("Welcome %s!\n", username()))
}

// print command line
func line(w io.Writer) {
	writeString(w, fmt.Sprintf("[%s:%s]$ ", username(), pwd()))
}

func username() string {
	return os.Getenv("USER")
}

// write string to writer
func writeString(w io.Writer, str ...string) {
	io.WriteString(w, strings.Join(str, " "))
}

// ############################################################### //
// Commands
// ############################################################### //

type InitCommand func(Command) func()

func CD(cmd Command) func() {
	return func() {
		if len(cmd.Args) == 0 {
			log.Println("cd: missing operand")

			return
		}

		cd(cmd.Args[0])
	}
}
func cd(path string) {
	if path == "" || path == "~" {
		path = os.Getenv("HOME")
	}

	if err := os.Chdir(path); err != nil {
		log.Println(err)
	}
}

func LS(cmd Command) func() {
	return func() {
		log.Println("args", cmd.Args)

		if len(cmd.Args) == 0 {
			cmd.Args = append(cmd.Args, ".")
		}

		writeString(cmd.w, ls(cmd.Args...))
	}
}
func ls(paths ...string) string {
	str := ""

	for i, v := range paths {
		dirEnties, err := os.ReadDir(v)
		if err != nil {
			str += err.Error() + "\n"

			continue
		}

		var info os.FileInfo
		for _, entry := range dirEnties {
			info, err = entry.Info()
			if err != nil {
				continue
			}

			str += fmt.Sprintf("%s  %s  %s\n", info.Mode(), info.ModTime().Format(time.Stamp), info.Name())
		}

		if len(paths) > i {
			str += "\n"
		}
	}

	return str
}

func Echo(cmd Command) func() {
	return func() {
		echo(cmd.w, cmd.Args)
	}
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

// PWD write current path
func PWD(cmd Command) func() {
	return func() {
		writeString(cmd.w, pwd())
	}
}
func pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)

		return ""
	}

	return dir
}

func PS(cmd Command) func() {
	return func() {
		writeString(cmd.w, ps())
	}
}
func ps() string {
	// тут лежат открытые процессы
	proc, err := os.Open("/proc")
	if err != nil {
		log.Println(err)

		return ""
	}

	// получаем информацию только о процессах
	// это имена директорий начинающиеся с номера процесса
	dirs, err := proc.Readdirnames(-1)
	if err != nil {
		log.Println(err)

		return ""
	}

	str := "PID\tTTY\tCMD\n"

	for _, v := range dirs {
		if v[0] < '0' || v[0] > '9' {
			continue
		}

		str += v + "\t"

		tty, _ := os.Readlink(fmt.Sprintf("/proc/%s/fd/0", v))
		str += tty + "\t"

		cmdline, _ := os.ReadFile(fmt.Sprintf("/proc/%s/cmdline", v))
		str += string(cmdline)
		str += "\n"
	}

	return str
}

func Kill(cmd Command) func() {
	return func() {
		if len(cmd.Args) == 0 {
			log.Println("kill: missing operand")

			return
		}

		pid, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			log.Println(err)

			return
		}

		kill(pid)
	}
}
func kill(pid int) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)

		return
	}

	if err = proc.Kill(); err != nil {
		log.Println(err)

		return
	}

	return
}

func Exec(cmd Command) func() {
	return func() {
		if len(cmd.Args) != 1 {
			log.Println("exec: wrong number of arguments")

			return
		}

		binary, err := exec.LookPath(cmd.Args[0])
		if err != nil {
			log.Println(err)

			return
		}

		args := append([]string{cmd.Name}, cmd.Args...)

		if err = syscall.Exec(binary, args, os.Environ()); err != nil {
			log.Println("exec error:", err)
		}
	}
}

func Fork(cmd Command) func() {
	return func() {
		ecmd := exec.Command(cmd.Name, cmd.Args...) //nolint:gosec

		log.Println("CMD ARGS", cmd, cmd.Args)

		ecmd.Stdout = cmd.w
		ecmd.Stderr = cmd.w
		ecmd.Stdin = cmd.r
		ecmd.Env = os.Environ()

		if err := ecmd.Run(); err != nil {
			log.Println("cmd run err:", err)

			return
		}

		log.Println(ecmd.Process.Pid)
	}
}

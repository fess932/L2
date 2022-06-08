package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.

towel.blinkenlights.nl // 23 port default

nc -l 1234 -> localhost:1234

*/

func main() {
	tf := flag.String("timeout", "10s", "timeout")
	flag.Parse()

	timeout, err := time.ParseDuration(*tf)
	if err != nil {
		log.Fatal(err)
	}

	addr := ""

	if flag.Arg(0) == "" {
		log.Fatal("no host")
	}

	addr += flag.Arg(0)

	if flag.Arg(1) == "" {
		addr += ":23"
	} else {
		addr += ":" + flag.Arg(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Println(gtelnet(ctx, cancel, addr))
}

func gtelnet(ctx context.Context, cancel context.CancelFunc, addr string) error {
	var dialer net.Dialer

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()

	// write
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			conn.Write([]byte(scanner.Text() + "\n"))
		}

		cancel()
	}()

	// read
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		cancel()
	}()

	<-ctx.Done()

	return nil
}

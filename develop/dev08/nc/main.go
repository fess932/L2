package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

// -u udp, default tcp
// hostname port

func main() {
	udp := flag.Bool("u", false, "udp")
	flag.Parse()

	network := "tcp"
	if *udp {
		network = "udp"
	}

	if len(flag.Args()) == 2 {
		nc(os.Stdin, network, flag.Args()[0], flag.Args()[1])
	} else {
		log.Println("nc: wrong number of arguments", flag.Args())
	}
}

func nc(r io.Reader, network, host, port string) {
	conn, err := net.Dial(network, host+":"+port)
	if err != nil {
		log.Println(err)

		return
	}
	defer conn.Close()

	// read from stdin, write to connection
	go func() {
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			io.WriteString(conn, scan.Text()+"\n")
		}
	}()

	// listen to sercve
	ch := make(chan struct{})
	// read from connection, write to stdout
	go func() {
		scan := bufio.NewScanner(conn)
		for scan.Scan() {
			io.WriteString(os.Stdout, scan.Text())
		}
		ch <- struct{}{}
	}()

	<-ch

	log.Println("nc: connection closed")
}

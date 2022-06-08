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
	log.Println(network)

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

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if _, err = io.WriteString(conn, scanner.Text()+"\n"); err != nil {
			log.Println(err)
		}
	}
}

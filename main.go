package main

import (
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	addrs := []string{":80", ":443"}
	ports, ok := os.LookupEnv("PORTS")
	if ok {
		addrs = strings.Split(ports, ",")
	}
	for _, addr := range addrs {
		go serve(addr)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func serve(address string) {
	log.Printf("listen address %v", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("listen address %v failed, %v\n", address, err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go response(conn)
	}
}

func response(conn net.Conn) {
	defer conn.Close()
	rip := conn.RemoteAddr().String()
	lip := conn.LocalAddr().String()
	defer log.Printf("%s -> %s disconnected", rip, lip)
	log.Printf("%s -> %s connected : ", rip, lip)
	buf := [1024]byte{}

	for {
		len, err := conn.Read(buf[:])
		if err != nil {
			return
		}

		_, err = conn.Write(buf[:len])
		if err != nil {
			return
		}
	}
}

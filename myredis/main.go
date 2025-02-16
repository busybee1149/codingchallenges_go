package main

import (
	"errors"
	"fmt"
	"io"
	"learngo/myredis/server"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "6379"
	TYPE = "tcp"
)


func main() {
	
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Started Redis server listening on %s\n", PORT)
	defer listener.Close()

	// Accept incoming connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
		    fmt.Println(err)
		    continue
		}
  
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	 }
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 32 * 1024)
	for {
	    n, err := conn.Read(buf)
	    if err != nil {
		   if errors.Is(err, io.EOF) {
			break
		   }
		   log.Println(err)
		   return
	    }
	    conn.Write([]byte(server.Serve(string(buf[:n]))))
	}
 }
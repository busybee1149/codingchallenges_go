package server

import (
	"errors"
	"fmt"
	"io"
	command "learngo/myredis/server/command"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "6379"
	TYPE = "tcp"
)

func Run() {
	listener, err := net.Listen(TYPE, HOST + ":" + PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Started Redis server listening on %s\n", PORT)
	defer listener.Close()

	connectionsLimit := make(chan int, 100)

	// Accept incoming connections and handle them
	for {
		connectionsLimit <- 1
		conn, err := listener.Accept()
		if err != nil {
		    fmt.Println(err)
		    continue
		}
		// Handle the connection in a new goroutine
		go safelyHandleConnection(conn, connectionsLimit)
	 }
}

func safelyHandleConnection(conn net.Conn, limit chan int) {
	defer func() {
		if err := recover(); err != nil {
		    log.Println("work failed:", err)
		}
	 }()
	 handleConnection(conn, limit)
}

func handleConnection(conn net.Conn, limit chan int) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
	    n, err := conn.Read(buf)
	    if err != nil {
		   if errors.Is(err, io.EOF) {
			break
		   }
		   log.Println(err)
		   return
	    }
	    conn.Write([]byte(command.Execute(string(buf[:n]))))
	}
	<- limit 
 }
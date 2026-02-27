package main

import (
	"fmt"
	"net"
	"os"
)

var Count int

func main() {
	args := os.Args
	if !(len(args) == 1 || len(args) == 2) {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	port := "8989"

	if len(args) != 1 {
		port = args[1]
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on the port :" + port)

	for {
		con, err := l.Accept()
		if err != nil {
			continue
		}

		mu.Lock()
		if Count < 10 {
			go HandlClient(con)
			Count++
			mu.Unlock()
		} else {
			mu.Unlock()
			con.Write([]byte("Sorry, we have 10 clients connexions in this room"))
			con.Close()

		}
	}
}

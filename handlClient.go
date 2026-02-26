package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	clients     = make(map[string]net.Conn)
	historiques []string
	mu          sync.Mutex
)

func HandlClient(con net.Conn, nbrClient int) {

	welcome, err := os.ReadFile("welcome.txt")
	if err != nil {
		fmt.Println(err)
	}
	con.Write(welcome)

	buf := make([]byte, 1024)
	n, _ := con.Read(buf)

	for string(buf[:n]) == "\n" {
		con.Write([]byte("[ENTER YOUR NAME]:"))
		n, _ = con.Read(buf)
		con.SetDeadline(time.Now().Add(2 * time.Second))
	}

	name := string(buf[:n-1])
	clients[name] = con

	now := time.Now()
	form := now.Format("[2006-01-02 15:04:05]")

	for n, c := range clients {
		if c != con {
			c.Write([]byte("\n" + name + " has joined our chat\n" + form + "[" + n + "]:"))
		}
	}

	for _, history := range historiques {
		for _, c := range clients {
			if c == con {
				c.Write([]byte(history))
			}
		}
	}

	for {

		ligne := form + "[" + name + "]:"

		con.Write([]byte(ligne))

		buff := make([]byte, 1024)
		nn, err := con.Read(buff)

		if err != nil {
			for n, c := range clients {
				if c != con {
					c.Write([]byte("\n" + name + " has left our chat...\n" + form + "[" + n + "]:"))
				}
			}
			Count--
			
			break

		}

		message := ligne + string(buff[:nn])

		mu.Lock()
		historiques = append(historiques, message)
		mu.Unlock()

		for n, c := range clients {
			if c != con {
				c.Write([]byte("\n" + message + form + "[" + n + "]:"))
			}
		}

	}

}

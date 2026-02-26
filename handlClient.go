package main

import (
	"bufio"
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
	form        string
)

func HandlClient(con net.Conn) {
	defer con.Close()

	buf := bufio.NewReader(con)

	welcome, err := os.ReadFile("welcome.txt")
	if err != nil {
		fmt.Println(err)
	}
	con.Write(welcome)

	nameInput, _ := buf.ReadString('\n')

	for {
		nameInput, _ = buf.ReadString('\n')

		if len(nameInput) == 0 {
			con.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}

		for _, r := range nameInput {
			if r < 32 || r > 126 {
				con.Write([]byte("[ENTER YOUR NAME]:"))
				continue
			}
		}
		mu.Lock()
		_, exists := clients[nameInput]
		mu.Unlock()
		if exists {
			con.Write([]byte("[NAME ALREADY EXISTS]\n"))
			con.Write([]byte("[ENTER YOUR NAME]:"))

			continue
		}

		break
	}

	name := nameInput[:len(nameInput)-1]
	mu.Lock()
	clients[name] = con
	mu.Unlock()

	now := time.Now()
	form = now.Format("[2006-01-02 15:04:05]")

	BroadCast(con, name+" has joined our chat\n")

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

		messageContent, err := buf.ReadString('\n')
		if err != nil {
			BroadCast(con, name+" has left our chat...\n")
			mu.Lock()
			Count--
			mu.Unlock()
			return
		}

		found := false
		for _, r := range messageContent {
			if (r < 32 || r > 126) && r != '\n' {
				found = true
				break
			}
		}
		if found {
			continue
		}
		message := ligne + messageContent
		mu.Lock()
		historiques = append(historiques, message)
		mu.Unlock()

		BroadCast(con, message)
	}
}

func BroadCast(con net.Conn, msg string) {
	for n, c := range clients {
		if c != con {
			c.Write([]byte("\n" + msg + form + "[" + n + "]:"))
		}
	}
}

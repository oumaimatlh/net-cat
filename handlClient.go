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
	formTime    string
	mu sync.Mutex
)

func HandlClient(con net.Conn) {
	defer con.Close()
	//Welcome
	welcome, err := os.ReadFile("welcome.txt")
	if err != nil {
		fmt.Println(err)
	}
	con.Write(welcome)
	//Validation => Name
	buf := bufio.NewReader(con)
	name := ""
	for {
		nameInput, _ := buf.ReadString('\n')
		if nameInput == "\n" {
			con.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}
		check := false
		for _, r := range nameInput {
			if (r < 32 || r > 126) && r != '\n' {
				con.Write([]byte("[ENTER YOUR NAME]:"))
				check = true
				break
			}
		}
		if check {
			continue
		}
		_, exists := clients[nameInput[:len(nameInput)-1]]
		if exists {
			con.Write([]byte("[NAME ALREADY EXISTS]\n"))
			con.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}
		name = nameInput[:len(nameInput)-1]
		break
	}
	//New Client
	mu.Lock()
	clients[name] = con
	mu.Unlock()

	//Time
	now := time.Now()
	formTime = now.Format("[2006-01-02 15:04:05]")

	//NewClient has Joined => CHAT
	BroadCast(con, name+" has joined our chat\n")

	//Show History for NewClient 
	for _, history := range historiques {
		for _, c := range clients {
			if c == con {
				c.Write([]byte(history))
			}
		}
	}
	
	for {
		ligne := formTime + "[" + name + "]:"
		con.Write([]byte(ligne))

		messageContent, err := buf.ReadString('\n')
		if err != nil {
			BroadCast(con, name+" has left our chat...\n")
			mu.Lock()
			Count--
			mu.Unlock()
			return
		}

		if messageContent=="\n"{
			continue
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
			c.Write([]byte("\n" + msg + formTime + "[" + n + "]:"))
		}
	}
}

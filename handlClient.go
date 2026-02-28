package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	clients     = make(map[string]net.Conn)
	historiques []string
	formTime    string
	mu          sync.Mutex
)

func HandlClient(con net.Conn) {
	defer con.Close()
	// Welcome
		welcome, err := os.ReadFile("welcome.txt")
		if err != nil {
			fmt.Println(err)
		}
		con.Write(welcome)
	// Validation => Name
		buf := bufio.NewReader(con)
		name := ""
		for {
			nameInput, err := buf.ReadString('\n')
			if err != nil {
				mu.Lock()
				Count--
				mu.Unlock()
				return
			}
			if nameInput == "\n" {
				con.Write([]byte("[ENTER YOUR NAME]:"))
				continue
			}
			check := false
			for _, r := range nameInput {
				if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) && r != '\n' {
					con.Write([]byte("[YOUR NAME SHOULD HAVE ALPHBA CARs \n ENTER YOUR NAME]:"))
					check = true
					break
				}
			}
			if check {
				continue
			}
			if len(nameInput) > 10 {
				con.Write([]byte("[YOUR NAME SHOULD HAVE MAX 10 CAR]\n[ENTER YOUR NAME]:"))
				continue

			}
			name = strings.TrimSuffix(nameInput, "\n")
			mu.Lock()
			_, exists := clients[name]
			mu.Unlock()
			if exists {
				con.Write([]byte("[NAME ALREADY EXISTS]\n[ENTER YOUR NAME]:"))
				continue
			}

			break
		}
	// New Client
		mu.Lock()
		clients[name] = con
		mu.Unlock()

	// Time
		now := time.Now()
		formTime = now.Format("[2006-01-02 15:04:05]")

	// NewClient has Joined => CHAT
		BroadCast(con, name+" has joined our chat\n")

	// Show History for NewClient
		mu.Lock()
		for _, history := range historiques {
			for _, c := range clients {
				if c == con {
					c.Write([]byte(history))
				}
			}
		}
		mu.Unlock()

	for {
		ligne := formTime + "[" + name + "]:"
		con.Write([]byte(ligne))

		messageContent, err := buf.ReadString('\n')
		if err != nil {
			BroadCast(con, name+" has left our chat...\n")

			mu.Lock()
			delete(clients, name)
			Count--
			mu.Unlock()
			return
		}

		if messageContent == "\n" {
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
	mu.Lock()
	for n, c := range clients {
		if c != con {
			c.Write([]byte("\n" + msg + formTime + "[" + n + "]:"))
		}
	}
	mu.Unlock()
}

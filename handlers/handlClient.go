package handlers

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	historiques []string
	mu          sync.Mutex
)
var clients = make(map[string]net.Conn)

func HandlClient(con net.Conn) {

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

	for {

		now := time.Now()
		form := now.Format("[2006-01-02 15:04:05]")
		ligne := form + "[" + name + "]:"

		con.Write([]byte(ligne))

		buff := make([]byte, 1024)
		nn, _ := con.Read(buff)

		message := ligne + string(buff[:nn])

		mu.Lock()
		historiques = append(historiques, message)
		mu.Unlock()

		for n, c := range clients {
			if c != con {
				c.Write([]byte("\n" + message + "\n" + form + "[" + n + "]:"))
			}
		}

		//fmt.Println(historiques)
		//fmt.Println("---------------------------------")

	}

}

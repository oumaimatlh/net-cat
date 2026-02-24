package handlers

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Chat struct {
	clients     map[string]*net.Conn
	historiques []string
	mu          sync.Mutex
}
var chat = Chat{
	clients: make(map[string]*net.Conn),
	historiques: []string{},
}
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
	

	chat.mu.Lock()
	chat.clients[name]=&con
	chat.mu.Unlock()	

	//write history
	for name, client :=range chat.clients{
		now := time.Now()
		form := now.Format("[2006-01-02 15:04:05]")
		ligne := form + "[" + name + "]:"

		client.con.Write([]byte(ligne))
		con.Read(buf)
		chat.mu.Lock()
		chat.historiques = append(chat.historiques, ligne+string(buf[:n]))
		chat.mu.Unlock()



	}


}

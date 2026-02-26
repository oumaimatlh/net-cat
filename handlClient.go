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

	//Welcome => Client
	welcome, err := os.ReadFile("welcome.txt")
	if err != nil {
		fmt.Println(err)
	}
	con.Write(welcome)

	buf := bufio.NewReader(con)

	//Gestion d'erreur => Entrez votre nom (Doublons d nom + Ascii pour le nom + nom est vide)
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

//Implementez map avec client courant 
	mu.Lock()
	clients[name] = con
	mu.Unlock()

	now := time.Now()
	formTime = now.Format("[2006-01-02 15:04:05]")

//Serveur sert a Distribué au clients =>  .. has joined our chat
	BroadCast(con, name+" has joined our chat\n")

//Serveur sert a Distribué l'historique s'ils existent pour le client courant 
	for _, history := range historiques {
		for _, c := range clients {
			if c == con {
				c.Write([]byte(history))
			}
		}
	}

//le Coeur de Serveur => hna  fach knjme3 les msg o nhethoum f historiques o serveru y distribuyehum les autres clients 
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


// o HADA BRODCAT li kayseft ge3 l msg dyl client courant l es autres clients 
func BroadCast(con net.Conn, msg string) {
	for n, c := range clients {
		if c != con {
			c.Write([]byte("\n" + msg + formTime + "[" + n + "]:"))
		}
	}
}

package handlers

import (
	"fmt"
	"net"
	"os"
	"time"
)

func HandlClient(con net.Conn) {
	// DATE // HEURE
	now := time.Now()
	form := now.Format("[2006-01-02 15:04:05]")

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

	ligne := form + "[" + name + "]:"

	Historiques := []string{}

	for string(buf[:n]) != "\n" {
		con.Write([]byte(ligne))
		con.Read(buf)
		Historiques = append(Historiques, ligne+string(buf[:n]))
		fmt.Println(Historiques)
	}
}

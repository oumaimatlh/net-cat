package main

import (
	"TEST/handlers"
	"fmt"
	"net"
	"os"
)



func main(){
	
	args:= os.Args
	if !(len(args) == 1 || len(args) == 2){
		fmt.Println("[USAGE]: ./TCPChat $port")
		return 
	}
	port := "8989"

	if len(args) != 1  {
		port = args[1]
	}


	l, err:= net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("Listening on the port :"+port)


	count := 0
	for {
		con, err := l.Accept()
		count++
		
		if err != nil {
			continue
		}
		go handlers.HandlClient(con)
		go fmt.Println(count)

	}


}
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	SocketPath := os.Args[1]
	if SocketPath == "help" {
		fmt.Println("Simple tool to send a message to an HTTP server via a socket file.")
		fmt.Println("Usage: httpSocket $SOCKET $SENDSTRING")
		os.Exit(1)
	}
	HTTPCall := os.Args[2]
	fmt.Println(HTTPCall)
	conn, err := net.Dial("unix", SocketPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not create connection to socket")
		os.Exit(2)
	}
	fmt.Fprintf(conn, "%s\r\n\r\n", HTTPCall)
	// response, err := bufio.NewReader(conn).ReadString('\n')
	//if err != nil {
	//	fmt.Println("I died.")
	//}

	output, err := ioutil.ReadAll(conn)
	finalOutput := fmt.Sprintf("%s", output)
	fmt.Print(finalOutput)
}

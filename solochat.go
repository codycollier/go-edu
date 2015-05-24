package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func logError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}

func handleConnection(connection net.Conn) {

	fmt.Printf("Incoming connection from: %s\n", connection.RemoteAddr())
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)

	writer.WriteString("Enter Nickname: ")
	writer.Flush()
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)
	prefix := fmt.Sprintf("%s> ", nick)

	for {
		line, _ := reader.ReadString('\n')
		writer.WriteString(prefix)
		writer.WriteString(line)
		writer.Flush()
	}

}

func main() {

	listener, err := net.Listen("tcp", ":2323")
	logError("listen error", err)
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		logError("Incoming connection error", err)
		defer connection.Close()
		go handleConnection(connection)
	}

}

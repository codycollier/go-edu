// partyline implements a basic chat server in the style of a party phone line
package main

import (
	"bufio"
	"net"
	"strings"
)

// chatUser is the structure and management of a single chat participant
type chatUser struct {
	reader *bufio.Reader
	writer *bufio.Writer
	nick   string
}

func (user *chatUser) setNick() {
	user.writer.WriteString("Enter Nickname: ")
	user.writer.Flush()
	nick, _ := user.reader.ReadString('\n')
	user.nick = strings.TrimSpace(nick)
}

func (user *chatUser) joinChat(partyline *partyLine) {

	partyline.addUser(user)

	for {
		line, _ := user.reader.ReadString('\n')
		// send to the partyline
		partyline.in <- line
	}

}

// partyLine manages the communication between all chatUser participants
type partyLine struct {
	users []*chatUser
	in    chan string
}

func (p *partyLine) addUser(user *chatUser) {
	p.users = append(p.users, user)
}

// handleConnection initializes new incoming users
func handleConnection(connection net.Conn, partyline *partyLine) {
	user := new(chatUser)
	user.reader = bufio.NewReader(connection)
	user.writer = bufio.NewWriter(connection)
	user.setNick()
	user.joinChat(partyline)
}

// Listen, accept connections, and connect them to the party line
func main() {

	partyline := new(partyLine)

	listener, _ := net.Listen("tcp", ":2323")
	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		defer connection.Close()
		go handleConnection(connection, partyline)
	}

}

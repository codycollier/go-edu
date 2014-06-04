package main

import (
	"log"
	"net"
)

func main() {

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	conn, err := net.ListenPacket("udp", ":0")
	defer conn.Close()

	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2323")
	if err != nil {
		log.Println("Error resolving remote address")
		log.Printf("Error: %s", err)
		log.Fatal()
	}

	packet := []byte("Test packet")
	//oob := make([]byte, 0)

	numBytes, err := conn.WriteTo(packet, remoteAddr)
	if err != nil {
		log.Println("Error sending udp packet")
		log.Printf("Err: %s", err)
		log.Fatal()
	}
	log.Printf("Sent %v bytes", numBytes)

}

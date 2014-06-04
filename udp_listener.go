package main

import (
	"log"
	"net"
)

func packet_handler(packet []byte, numBytesPacket int, fromAddr net.Addr) {
	log.Printf("Incoming packet from %s (%v bytes) in: %s\n", fromAddr, numBytesPacket, string(packet))
}

func main() {

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	listenAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:2323")
	srv, err := net.ListenUDP("udp", listenAddr)
	defer srv.Close()
	if err != nil {
		log.Println("Error starting service")
		log.Printf("Error: %v", err)
		log.Fatal()
	}

	packet := make([]byte, 64*1024)
	oob := make([]byte, 64*1024)

	for {

		numBytesPacket, numBytesOOB, flags, fromAddr, err := srv.ReadMsgUDP(packet, oob)
		if err != nil {
			log.Println("Error accepting incoming packet")
			log.Printf("Error: %v", err)
		}

		_ = numBytesOOB
		_ = flags

		if numBytesPacket > 0 {
			go packet_handler(packet, numBytesPacket, fromAddr)
		}
	}

}

// boxy is some testing with the NaCl libs in go
package main

import (
	"code.google.com/p/go.crypto/nacl/box"
	"crypto/rand"
	"fmt"
)

func main() {

	pubkey, prvkey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf(" public key:\n%s\n", pubkey)
	fmt.Printf("private key:\n%s\n", prvkey)

}

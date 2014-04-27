// boxy - working with the NaCl libs in go
package main

import (
	"code.google.com/p/go.crypto/nacl/box"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {

	pubkey, prvkey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf(" public key (hex):\n%x\n", *pubkey)
	fmt.Printf("private key (hex):\n%x\n", *prvkey)

	hash256 := sha256.New()

	hash256.Write(pubkey[:])
	fmt.Printf(" public key (sha256 hash):\n%x\n", hash256.Sum(nil))

	hash256.Reset()
	hash256.Write(prvkey[:])
	fmt.Printf("private key (sha256 hash):\n%x\n", hash256.Sum(nil))

}

// boxy - working with the NaCl libs in go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/nacl/box"
)

func main() {

	for i := 0; i < 100000; i++ {
		pubkey, prvkey, err := box.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		fmt.Printf("        (hex)  public key: %x\n", *pubkey)
		fmt.Printf("        (hex) private key: %x\n", *prvkey)

		hash256 := sha256.New()

		hash256.Write(pubkey[:])
		fmt.Printf("(sha256 hash)  public key: %x\n", hash256.Sum(nil))

		hash256.Reset()
		hash256.Write(prvkey[:])
		fmt.Printf("(sha256 hash) private key: %x\n", hash256.Sum(nil))
	}

}

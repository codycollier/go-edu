package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	csid_string := "3a"
	fmt.Printf("csid_string: %s\n", csid_string)

	csid_byte, _ := hex.DecodeString(csid_string)
	fmt.Printf("len(csid_byte): %s\n", len(csid_byte))

	csid_string_again := hex.EncodeToString(csid_byte)
	fmt.Printf("csid_string: %s\n", csid_string_again)
}

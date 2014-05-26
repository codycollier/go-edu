package main

import (
	"encoding/json"
	"fmt"
)

type message1 struct {
	name  string
	mtype string
}

type message2 struct {
	Name  string
	Mtype string "Type"
}

type message3 struct {
	Name string
	Type string
}

func main() {

	var j1 []byte
	var j1m1 message1
	var j1m2 message2
	var j1m3 message3

	j1 = []byte(`{"name":"namej1", "type":"typej1"}`)
	fmt.Println("The initial json:")
	fmt.Printf("j1: %s\n", j1)

	json.Unmarshal(j1, &j1m1)
	json.Unmarshal(j1, &j1m2)
	json.Unmarshal(j1, &j1m3)
	fmt.Println("")
	fmt.Println("The json unmarshaled in to structs:")
	fmt.Printf("j1 into m1: %s\n", j1m1)
	fmt.Printf("j1 into m2: %s\n", j1m2)
	fmt.Printf("j1 into m3: %s\n", j1m3)

	fmt.Println("")
	fmt.Println("The desired struct:")
	m0 := message2{"namej1", "typej1"}
	fmt.Printf("          : %s\n", m0)

	fmt.Println("")
}

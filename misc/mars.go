package main

import (
	"encoding/json"
	"fmt"
)

//
// Unmarshal
//

type message1 struct {
	name  string
	mtype string
}

type message2 struct {
	Name   string `json:"name"`
	Chtype string `json:"Type"`
}

type message3 struct {
	Name string
	Type string
}

//
// Marshal
//

type messageA struct {
	name   string `json:"name"`
	chtype string `json:"type"`
}

type messageB struct {
	Name   string `json:"name"`
	ChType string `json:"type"`
}

func main() {

	var j1 []byte
	var j1m1 message1
	var j1m2 message2
	var j1m3 message3
	var m0 message3

	fmt.Println("")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(":: unmarshal")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("")

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
	m0 = message3{"namej1", "typej1"}
	fmt.Printf("          : %s\n", m0)

	fmt.Println("")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(":: marshal")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("")

	var jA []byte
	var mA messageA
	var mB messageB

	mA = messageA{"namemA", "typemA"}
	mB = messageB{"namemB", "typemB"}
	fmt.Println("The initial struct...")
	fmt.Printf("mA: %s\n", mA)
	fmt.Printf("mB: %s\n", mB)

	jA, _ = json.Marshal(mA)
	jA, _ = json.Marshal(mB)
	fmt.Println("")
	fmt.Println("The struct marshaled in to json:")
	fmt.Printf("mA into json: %s\n", jA)
	fmt.Printf("mB into json: %s\n", jA)

	fmt.Println("")
	fmt.Println("")
}

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
	Mtype string
}

func main() {
	var err error
	var j1 []byte
	var j2 []byte
	var m1 message1
	var m2 message2

	j1 = []byte(`{"name":"green", "type":"color"}`)
	j2 = []byte(`{"name":"green", "type":"color"}`)
	fmt.Printf("j1: %s\n", j1)
	fmt.Printf("j2: %s\n", j2)

	err = json.Unmarshal(j1, &m1)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Printf("m1: %s\n", m1)

	err = json.Unmarshal(j2, &m2)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Printf("m2: %s\n", m2)

}

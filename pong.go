//
// pong - exploring goroutines and channels
//
// reference / inspiration:
// http://www.youtube.com/watch?v=QDDwwePbDtw (Sameer Ajmani)
//

// pong is an exploration of goroutines and channels
package main

import (
	"fmt"
	"time"
)

type Ball struct{ hitcount int }
type Table chan *Ball

func player(label string, table Table) {
	for {
		ball := <-table
		ball.hitcount++
		fmt.Printf("player:: %s :: hit!\n", label)
		table <- ball
	}
}

func main() {
	ball := new(Ball)
	table := make(Table)
	go player("one", table)
	go player("two", table)
	table <- ball

	time.Sleep(time.Millisecond * 300)

	// Grab the ball off the table
	b := <-table
	fmt.Printf("hit count :: %d\n", b.hitcount)

	panic("Panic to show info")

}

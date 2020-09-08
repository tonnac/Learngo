package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"zxczxc", "aqwew"}
	for index, person := range people {

		switch d := index % 2; d {
		case 0:
			go isSexy(person, true, c)
		case 1:
			go isSexy(person, false, c)
		default:
			go isSexy(person, false, c)
		}
	}
	// result := <-c
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, re bool, c chan bool) {
	fmt.Printf("start %s\n", person)
	time.Sleep(time.Second * 5)
	fmt.Printf("end %s\n", person)
	c <- re
}

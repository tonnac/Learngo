package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("qwe")
	sexyCount("das")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second)
	}
}

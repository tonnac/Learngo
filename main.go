package main

import (
	"fmt"
	"strings"
)

//Naked
func letAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := letAndUpper("ewwls")
	fmt.Println(totalLength, upperName)

}

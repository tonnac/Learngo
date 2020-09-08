package main

import (
	"fmt"
	"strings"
)

func letAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := letAndUpper("ewwls")
	totalLength2, _ := letAndUpper("ewqwewqwls")
	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2)

	repeatMe("qqqq", "wwww", "eeee", "rrrr", "ttttt", "rrrrr", "tttt")
}

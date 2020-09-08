package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Kimchi", "Ramen"}
	fsdf := person{name: "dddd", age: 18, favFood: favFood}
	fmt.Println(fsdf.name)
}

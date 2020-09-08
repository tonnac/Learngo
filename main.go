package main

import "fmt"

func main() {
	names := []string{"qweq", "adasd", "dsfdsfd"}
	names0 := append(names, "ekrjer")
	names1 := append(names0, "pwerewr")
	names2 := append(names1, "mmkwe")
	fmt.Println(names)
	fmt.Println(names0)
	fmt.Println(names1)
	fmt.Println(names2)
}

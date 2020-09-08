package main

import "fmt"

func main() {
	// lkes := map[string]string{"name": "nico", "age": "12"}
	kfd := [5]map[string]string{
		{"name": "qweqwe", "age": "32"},
		{"name0": "asdasd", "age0": "42"},
		{"name1": "zxczxc", "age1": "52"},
		{"name2": "rtyrty", "age2": "62"},
		{"name3": "fghfgh", "age3": "72"}}
	// for _, value := range lkes {
	// 	fmt.Println(value)
	// }

	for _, map0 := range kfd {
		for key, value := range map0 {
			fmt.Println(key, value)
		}
	}

	// fmt.Println(kfd)
	// fmt.Println(lkes)
}

package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["jose"] = 14
	m["nicolas"] = 28

	fmt.Println(m)

	// recorrer maps

	for i, v := range m {
		println(i, v)
	}

	//encontrar un valor

	value, _ := m["jos"]
	fmt.Println(value)
	fmt.Printf("2, %T \n", 2)

	m1 := make(map[string]int)
	m1["a"] = 8

	fmt.Println(m1)

}

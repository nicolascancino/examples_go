package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	c := make(chan string)
	start := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidores(servidor, c)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-c)
	}
	end := time.Since(start)
	fmt.Println("Tiempo de ejecucion: ", end)
}

func revisarServidores(servidor string, channel chan<- string) {
	_, err := http.Get(servidor)

	if err != nil {
		channel <- servidor + "no se encuentra disponible"
	} else {
		channel <- servidor + "esta funcionado normalmente"
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int, 2)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomethingV2(i, &wg, c)
	}

	wg.Wait()

}

func doSomethingV2(i int, wg *sync.WaitGroup, c chan int) {

	defer wg.Done()
	fmt.Println("Inicio", i)
	time.Sleep(1 * time.Second)
	fmt.Println("Termino")
	<-c
}

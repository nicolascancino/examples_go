package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func message(text string, c chan string) {
	c <- text
}
func main() {

	c := make(chan string, 1)

	fmt.Println("hello")

	go say("bye", c)

	fmt.Println(<-c)

	// second example

	c2 := make(chan string, 2)

	c2 <- "Mensaje 1"
	c2 <- "Mensaje 2"

	fmt.Println(len(c2), cap(c2))

	// close the channel
	close(c2)

	//range for a channel
	for menssage := range c2 {
		fmt.Println(menssage)
	}

	// Select

	email1 := make(chan string)
	email2 := make(chan string)

	go message("message1", email1)
	go message("message2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
		close(email1)
		close(email2)
	}

}

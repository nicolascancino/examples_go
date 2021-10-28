package main

import (
	"fmt"
)

func main() {

	server := NewServer(":8080")
	server.healthcheck("/healthcheck", Healthcheck())
	server.handle("/", "GET", server.AddMiddleware(HandleRoot))
	server.handle("/api", "POST", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.handle("/create", "POST", PostRequest)
	fmt.Println(server.listen())
}

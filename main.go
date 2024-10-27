package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	server := NewApiServer(":8080")

	server.Run()
}

package main

import "log"

func main() {
	server := NewServer()
	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	web "cyoa/web"
	"log"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatal("Please provide a supported command to run")
	}

	cmd := os.Args[1]

	if cmd == "web" {
		web.Parser.Parse(os.Args[2:])
		web.Run()
	}
}

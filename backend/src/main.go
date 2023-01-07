package main

import (
	"log"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Println(rest.RunAPI("localhost:8000"))
}

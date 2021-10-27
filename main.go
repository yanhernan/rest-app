package main

import (
	"log"

	"github.com/yanhernan/rest-app/routes"
)

func main() {
	router := routes.Default()
	log.Println("Started the server")
	err := router.Start()
	if err != nil {
		log.Fatalln(err)
	}

}

package main

import (
	"fmt"
	"log"
	"tung.gallery/internal/api"
)

func main() {
	r := api.NewRouter()
	api.Initialize(r)
	fmt.Println("start run server")
	if err := r.Engine.Run(":5000"); err != nil {
		log.Fatalln(err.Error())
	}
}

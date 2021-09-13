package main

import (
	"fmt"
	"tung.gallery/internal/api"
)

func main() {
	r := api.NewRouter()
	api.Initialize(r)
	fmt.Println("start run server")
	r.Engine.Run(":5000")
}

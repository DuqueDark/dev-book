package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	fmt.Println("Rodando API!")
	log.Fatal(http.ListenAndServe(":5000", r))
}

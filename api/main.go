package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API running...")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}

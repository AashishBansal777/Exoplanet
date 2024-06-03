package main

import (
	"ExoPlanet/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("server sucessfully listen at : 8080")
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kdkumawat/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing at port 4000 ...")
}

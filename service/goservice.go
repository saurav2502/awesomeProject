package service

import (
	"fmt"
	"log"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {
	log.Print("entered to Handle function")
	fmt.Fprintf(writer, "App is up")
}

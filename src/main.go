package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Chirag-Nayak/go-basics/web-service/handlers"
)

func main() {

	l := log.New(os.Stdout, "Product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}

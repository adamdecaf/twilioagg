package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"

	"github.com/adamdecaf/twilioagg/routes"
)

const (
	defaultHttpPort = 8080
)

var (
	port = flag.Int("port", defaultHttpPort, "port to bind http server on")
)

func main() {
	flag.Parse()
	log.Println("starting twilioagg")

	err := startHttpServer(*port)
	if err != nil {
		log.Fatalf("error in starting http server - %s", err)
	}
}

func startHttpServer(port int) error {
	// register handlers
	http.HandleFunc("/ping", routes.Ping)

	listen := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listen, nil)
	return err
}

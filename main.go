package main

import (
	"flag"
	"fmt"
	"github.com/rmukubvu/amakhosi/handlers"
	"log"
	"net/http"
)

var port = flag.Int("p", 8000, "port number")

func main() {
	flag.Parse()
	r := handlers.InitRouter()
	//convert to port format
	sPort := fmt.Sprintf(":%d", *port)
	//show on stdout
	fmt.Printf("Connecting to port [%s]", sPort)
	log.Fatal(http.ListenAndServe(sPort, r))
}

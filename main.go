package main

import (
	"flag"
	"fmt"
	"github.com/rmukubvu/amakhosi/handlers"
	"github.com/rmukubvu/amakhosi/store"
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
	err := http.ListenAndServe(sPort, r)
	if err != nil {
		//close db
		store.CloseDb()
		//log error
		log.Fatal(err)
	}
}

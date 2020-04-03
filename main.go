package main

import "fmt"

//var port = flag.Int("p", 8000, "port number")

func main() {
	/*flag.Parse()
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
	}*/
	test(1, 2, 3, "hello")
	test()
}

func test(args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf("found some args %+v", args)
	} else {
		fmt.Printf("no args available")
	}
}

package main

import (
	"fmt"
	"github.com/rmukubvu/amakhosi/store"
)

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
	test("27719084111", "7903245789130", 1)

}

func test(args ...interface{}) {
	const query = "INSERT INTO accounts (account_number,personal_identifier,is_active,created_date) VALUES (?,?,?,now())"
	a, err := store.Insert(query, args)
	if err != nil {
		fmt.Errorf("insert error %v", err)
	}
	fmt.Println(a)
}

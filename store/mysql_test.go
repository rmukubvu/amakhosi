package store

import (
	"fmt"
	"strings"
	"testing"
)

//TODO: look for a better place to put the tests
//cant be copying the config file to every folder

/*
func TestInsert(t *testing.T) {
	var want interface{} = nil
	m := &model.Account{
		AccountNumber: "27721681688",
		Identifier:    "78968963366",
		IsActive:      false,
		CreatedDate:   "2020-04-04",
	}

	if got := Insert("accounts",m) ; got != want {
		t.Fatalf("Expected nil , got %v", got)
	}
}
*/

func TestFetch(t *testing.T) {
	want := "27719084111"
	mp, err := Fetch("select * from accounts where id = ?", 5)
	if err != nil {
		t.Fatalf("failed to fetch data")
	}
	res := mp["account_number"]
	sres := fmt.Sprintf("%s", res)
	if strings.Compare(sres, want) != 0 {
		t.Fatalf("res is %s", sres)
	}
	/*m := &model.Account{
		AccountNumber: "27719084111",
		Identifier:    "7896366",
		IsActive:      false,
		CreatedDate:   "2020-04-04",
	}
	md,err := db.BindModel("accounts",m)
	if err != nil {
		t.Fatalf("failed to bind to model")
	}

	id := 5
	q := md.Select("*").Where(md.C("id").Eq(5))
	//q := md.Select(md,id)
	var res []model.Account
	err = db.Select(&res,q)
	if strings.Compare(res[0].AccountNumber,want) != 0 {
		t.Fatalf("no record exists with this id %d , res is %s",id,res[0].AccountNumber)
	}*/
}

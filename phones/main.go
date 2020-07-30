package main

import (
	"fmt"
	"gophercises/phones/normalize"
	"gophercises/phones/phonesdb"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "icmadmin"
	password = "admin"
	dbname   = "phones_db"
)

func main() {
	phonesdb.Init(false)

	phones, err := phonesdb.AllPhones()
	if err != nil {
		fmt.Println("Error getting all phones:", err)
	}
	for _, phone := range phones {
		fmt.Printf("The phone \"%s\" was normalized to \"%s\"\n", phone, normalize.Normalize(phone))
	}
}

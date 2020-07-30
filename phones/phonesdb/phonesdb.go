package phonesdb

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "icmadmin"
	password = "admin"
	dbname   = "phones_db"
)

var psqlinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
var db *sql.DB

func Init() {
	fmt.Println("Init phonesdb function")
	var err error
	db, err = sql.Open("postgres", psqlinfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	createDatabase()
	insertInitialData()
}

func truncatePhonesTable() {
	sqlStatement := "TRUNCATE TABLE phone_numbers;"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func createDatabase() {
	sqlStatement := `
	CREATE TABLE phone_numbers (
		id SERIAL PRIMARY KEY,
		phone TEXT
	  );
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func insertInitialData() {
	sqlStatements := `
		INSERT INTO phone_numbers (phone)
		VALUES ('1234567890');
		INSERT INTO phone_numbers (phone)
		VALUES ('123 456 7891');
		INSERT INTO phone_numbers (phone)
		VALUES ('(123) 456 7892');
		INSERT INTO phone_numbers (phone)
		VALUES ('(123) 456-7893');
		INSERT INTO phone_numbers (phone)
		VALUES ('123-456-7894');
	`
	_, err := db.Exec(sqlStatements)
	if err != nil {
		panic(err)
	}
}

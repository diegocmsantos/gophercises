package postgresconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwe123*"
	dbname   = "urlshortdb"
)

// GetConn returns postgresql connection
func Getconn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database ERORRRRRRRR")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil
}

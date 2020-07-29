package main

import (
	"fmt"
	"gophercises/todo/cmd"
	"gophercises/todo/db"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	homeDirectory, _ := homedir.Dir()
	dbPath := filepath.Join(homeDirectory, "todos.db")
	must(db.Init(dbPath, "todos"))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

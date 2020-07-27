package manager

import (
	"fmt"
	"gophercises/todo/db"
	"strconv"
)

const bucketName = "todos"

type Todo struct {
	Description string
	Done        bool
}

// Create will store the new Todo and return it
func Create(todo *Todo) (*Todo, error) {
	conn, err := db.Open()
	defer conn.Close()
	if err != nil {
		return nil, fmt.Errorf("list: error getting the database connection: %q", err)
	}

	if err := db.Update(conn, bucketName, todo.Description, strconv.FormatBool(todo.Done)); err != nil {
		return nil, fmt.Errorf("create: error creating %v: %q", todo, err)
	}
	return todo, nil
}

// List it lists all not done todos
func List() ([]Todo, error) {
	conn, err := db.Open()
	defer conn.Close()
	if err != nil {
		return nil, fmt.Errorf("list: error getting the database connection: %q", err)
	}

	var todos []Todo
	if conn != nil {
		todosFound, err := db.ReadAll(conn, bucketName)
		if err != nil {
			return nil, fmt.Errorf("list: error listing todos: %q", err)
		}

		for key, value := range todosFound {
			done, _ := strconv.ParseBool(value)
			if !done {
				todos = append(todos, Todo{Description: key, Done: done})
			}
		}
	}

	return todos, nil
}

// MarkAsDone will mark a specific todo as done
func MarkAsDone(todo *Todo) error {
	conn, err := db.Open()
	defer conn.Close()
	if err != nil {
		return fmt.Errorf("markAsDone: error getting the database connection: %q", err)
	}

	if conn != nil {
		if err := db.Update(conn, bucketName, todo.Description, strconv.FormatBool(todo.Done)); err != nil {
			return fmt.Errorf("markAsDone: error making todo \"%+v\" as done: %q", todo, err)
		}
	}

	return nil

}

package manager

import (
	"fmt"
	"gophercises/todo/db"
)

const bucketName = "todos"

type Todo struct {
	Key   int
	Value string
}

// Create will store the new Todo and return it
func Create(todo *Todo) (*Todo, error) {

	_, err := db.Update(bucketName, todo.Value)
	if err != nil {
		return nil, fmt.Errorf("create: error creating %v: %q", todo, err)
	}
	return todo, nil
}

// List it lists all not done todos
func List() ([]Todo, error) {

	var todos []Todo
	todosFound, err := db.ReadAll(bucketName)
	if err != nil {
		return nil, fmt.Errorf("list: error listing todos: %q", err)
	}

	for key, value := range todosFound {
		todos = append(todos, Todo{Key: key, Value: value})
	}

	return todos, nil
}

// MarkAsDone will mark a specific todo as done
func MarkAsDone(key int) error {
	return db.Delete(bucketName, key)
}

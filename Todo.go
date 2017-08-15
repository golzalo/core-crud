package main

import "fmt"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed int    `json:"completed" binding:"required"` //TODO investigar porq me exige tener 1 como valor y no cero
}

var memory = map[int]*Todo{}
var counter int = 0

func add(todo *Todo) int {
	counter++
	todo.ID = counter
	memory[counter] = todo
	return counter
}

func remove(id int) {
	delete(memory, id)
}

func get(id int) *Todo {
	return memory[id]
}

func getAll() []*Todo {
	v := make([]*Todo, 0, len(memory))
	for _, value := range memory {
		v = append(v, value)
	}
	return v
}

func update(id int, todo *Todo) {
	todo.ID = id
	memory[id] = todo
}

func PrintAll() {
	for key, value := range memory {
		fmt.Println("GON", key, value)
	}
	fmt.Println("")
}

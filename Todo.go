package main

import (
	"fmt"
	"sort"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed int    `json:"completed" default:"0"`
}

type TodoSortWrapper []*Todo

func (c TodoSortWrapper) Len() int           { return len(c) }
func (c TodoSortWrapper) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c TodoSortWrapper) Less(i, j int) bool { return c[i].ID > c[j].ID }

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
	sort.Sort(TodoSortWrapper(v))
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

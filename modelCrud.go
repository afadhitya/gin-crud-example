package main

import "github.com/jinzhu/gorm"

type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

type transformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

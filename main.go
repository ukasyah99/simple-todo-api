package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ukasyah99.dev/simple-todo-api/db"
	"ukasyah99.dev/simple-todo-api/handler"
)

func init() {
	godotenv.Load()
	db.Init()
	db.Migrate()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todo", handler.ListTodo)
	r.POST("/todo", handler.AddTodo)
	r.PUT("/todo/:id", handler.UpdateTodo)
	r.DELETE("/todo/:id", handler.DeleteTodo)

	r.Run()
}

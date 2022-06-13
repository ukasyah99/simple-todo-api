package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ukasyah99.dev/simple-todo-api/db"
)

func ListTodo(c *gin.Context) {
	var todos []db.Todo
	if err := db.DB.Order("id DESC").Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

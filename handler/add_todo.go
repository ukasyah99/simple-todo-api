package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ukasyah99.dev/simple-todo-api/db"
)

func AddTodo(c *gin.Context) {
	var form db.Todo
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": form,
	})
}

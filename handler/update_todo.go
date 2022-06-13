package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ukasyah99.dev/simple-todo-api/db"
)

type updateTodoForm struct {
	Title string `form:"title"`
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var form updateTodoForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&db.Todo{}).Where("id = ?", id).Update("title", form.Title)

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully",
	})
}

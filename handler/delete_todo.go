package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ukasyah99.dev/simple-todo-api/db"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := db.DB.Where("id = ?", id).Unscoped().Delete(&db.Todo{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}

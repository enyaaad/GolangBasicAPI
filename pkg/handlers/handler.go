package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var notes = []note{
	{ID: "1", Title: "Buy food", Description: "Go shop and buy something"},
	{ID: "2", Title: "Do workout", Description: "Get six press cubes"},
	{ID: "3", Title: "Kill Kenny", Description: "He must die"},
}

func GetAllNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

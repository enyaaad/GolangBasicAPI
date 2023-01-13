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
func GetNoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range notes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
func PostNote(c *gin.Context) {
	var newNotes note

	if err := c.BindJSON(&newNotes); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	notes = append(notes, newNotes)
	c.IndentedJSON(http.StatusCreated, newNotes)
}

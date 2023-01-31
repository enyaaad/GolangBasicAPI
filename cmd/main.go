package main

import (
	H "github.com/enyaaad/trainig/pkg/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file")
	}
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"192.168.1.2"}); err != nil {
		return nil
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/notes", H.GetAllNotes)
	r.GET("/notes/:id", H.GetNoteByID)
	r.POST("/notes", H.PostNote)

	return r
}
func main() {

	if err := initConfig(); err != nil {
		log.Printf("error init config.yml: %s", err.Error())
	}
	mode, exist := os.LookupEnv("GIN_MODE")

	if exist {
		gin.SetMode(mode)
		log.Printf(mode)
	}

	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

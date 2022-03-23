package main

import (
	// "net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/jason0x43/go-toggl"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:4000")
}

func getCurrentTogglEntry() toggl.TimeEntry {
	t := toggl.OpenSession(os.Getenv("TOGGL_TOKEN"))
	entry, err := t.GetCurrentTimeEntry()

	if err != nil {
	}

	return entry
}

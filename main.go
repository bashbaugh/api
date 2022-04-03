package main

import (
	"os"
	"strconv"

	"github.com/bashbaugh/api/lib"
	"github.com/gin-gonic/gin"
	"github.com/jason0x43/go-toggl"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	config := lib.LoadConfig()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/now", func(c *gin.Context) {
		entry := getCurrentTogglEntry()
		name := config.Toggl.Projects[strconv.Itoa(entry.Pid)]

		isTracking := entry.ID != 0

		c.JSON(200, gin.H{
			"trackingTime": isTracking,
			"activityName": name,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	r.Run("0.0.0.0:" + port)
}

func getCurrentTogglEntry() toggl.TimeEntry {
	t := toggl.OpenSession(os.Getenv("TOGGL_TOKEN"))
	entry, err := t.GetCurrentTimeEntry()

	if err != nil {
	}

	return entry
}

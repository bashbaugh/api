package controllers

import (
	"os"
	"strconv"

	"github.com/bashbaugh/api/lib"
	"github.com/gin-gonic/gin"
	"github.com/jason0x43/go-toggl"
)

var config = lib.LoadConfig()

func GetCurrentActivity(c *gin.Context) {
	entry := getCurrentTogglEntry()
	proj := config.Toggl.Projects[strconv.Itoa(entry.Pid)]

	isTracking := entry.ID != 0

	c.JSON(200, gin.H{
		"trackingTime":        isTracking,
		"activityName":        proj.Name,
		"activityDescription": proj.Description,
	})
}

func getCurrentTogglEntry() toggl.TimeEntry {
	t := toggl.OpenSession(os.Getenv("TOGGL_TOKEN"))
	entry, err := t.GetCurrentTimeEntry()

	if err != nil {
	}

	return entry
}

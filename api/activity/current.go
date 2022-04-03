package currentActivityHandler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/bashbaugh/api/lib"
	"github.com/jason0x43/go-toggl"
)

type HandlerResult struct {
	IsTracking bool   `json:"trackingTime"`
	Name       string `json:"activityName"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	config := lib.LoadConfig()

	entry := getCurrentTogglEntry()
	Name := config.Toggl.Projects[strconv.Itoa(entry.Pid)]

	IsTracking := entry.ID != 0

	json.NewEncoder(w).Encode(HandlerResult{IsTracking, Name})
}

func getCurrentTogglEntry() toggl.TimeEntry {
	t := toggl.OpenSession(os.Getenv("TOGGL_TOKEN"))
	entry, err := t.GetCurrentTimeEntry()

	if err != nil {
	}

	return entry
}

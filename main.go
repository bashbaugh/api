package main

import (
	"os"

	"github.com/jason0x43/go-toggl"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
}

func getCurrentTogglEntry() toggl.TimeEntry {
	t := toggl.OpenSession(os.Getenv("TOGGL_TOKEN"))
	entry, err := t.GetCurrentTimeEntry()

	if err != nil {
	}

	return entry
}

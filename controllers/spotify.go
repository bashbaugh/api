package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SpotifyCurrentSongResponse struct {
	IsPlaying bool `json:"is_playing"`
	Item      struct {
		Name  string `json:"name"`
		Album struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"album"`
	} `json:"item"`
}

func GetCurrentSong(c *gin.Context) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/currently-playing", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SPOTIFY_OAUTH_TOKEN"))

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		c.JSON(500, gin.H{
			"error": "Couldn't get current song",
		})
		return
	}

	decoder := json.NewDecoder(response.Body)
	var data SpotifyCurrentSongResponse
	decoder.Decode(&data)

	fmt.Println(data)

	c.JSON(200, gin.H{
		"isPlaying":  data.IsPlaying,
		"songName":   data.Item.Name,
		"artistName": data.Item.Album.Artists[0].Name,
	})
}

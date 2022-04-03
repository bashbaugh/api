package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

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
	req.Header.Set("Authorization", "Bearer "+GetSpotifyAccessToken())

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

	c.JSON(200, gin.H{
		"isPlaying":  data.IsPlaying,
		"songName":   data.Item.Name,
		"artistName": data.Item.Album.Artists[0].Name,
	})
}

func GetSpotifyAccessToken() string {
	client := &http.Client{}

	formData := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {os.Getenv("SPOTIFY_REFRESH_TOKEN")},
	}

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(formData.Encode()))
	req.Header.Set("Authorization", "Basic "+os.Getenv("SPOTIFY_AUTH_HEADER_VAL_BASE64"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	decoder := json.NewDecoder(response.Body)
	var tokenData struct {
		AccessToken string `json:"access_token"`
	}
	decoder.Decode(&tokenData)

	return tokenData.AccessToken
}

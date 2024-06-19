package controllers

import (
	"fmt"
	"net/http"

	"github.com/admiralyeoj/anime-organizer/internal/sonarr"
)

// Home Handler is the handler for our `/` url
func GetSeriesHandler(w http.ResponseWriter, _ *http.Request) {
	cfg := sonarr.NewClient(5)

	series, err := cfg.GetSeries()

	if err != nil {
		fmt.Println(err)
	}

	for _, value := range series {
		fmt.Println(value)
		break
	}
}

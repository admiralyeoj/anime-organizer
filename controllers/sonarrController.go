package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/admiralyeoj/anime-organizer/internal/sonarr"
)

func GetSeriesHandler(w http.ResponseWriter, _ *http.Request) {
	cfg := sonarr.NewClient(5)

	series, err := cfg.GetSeries()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(series))

	// Create response from message
	resp := NewJSONSuccessResponse(http.StatusOK, series)

	// Write our response
	if err := resp.WriteResponse(w); err != nil {
		log.Fatalf("failed to write a response: %s", err.Error())
	}
}

func GetEpisodesHandler(w http.ResponseWriter, r *http.Request) {
	cfg := sonarr.NewClient(5)

	query := r.URL.Query()
	seriesId := query.Get("seriesId")
	seasonNumber := query.Get("seasonNumber")

	if seriesId == "" {
		// Create response from message
		resp := NewJSONErrorResponse(http.StatusBadRequest, "")

		// Write our response
		if err := resp.WriteResponse(w); err != nil {
			log.Fatalf("failed to write a response: %s", err.Error())
		}
		return
	}

	episodes, err := cfg.GetEpisodes(&seriesId, &seasonNumber)

	if err != nil {
		fmt.Println(err)
	}

	// Create response from message
	resp := NewJSONSuccessResponse(http.StatusOK, episodes)

	// Write our response
	if err := resp.WriteResponse(w); err != nil {
		log.Fatalf("failed to write a response: %s", err.Error())
	}
}

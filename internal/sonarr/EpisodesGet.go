package sonarr

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/admiralyeoj/anime-organizer/internal/sonarr/models"
)

func (c *Client) GetEpisodes(seriesId, seasonNumber *string) ([]models.Episode, error) {
	url := c.baseUrl + "/v3/episode?"

	if seriesId == nil || *seriesId == "" {
		return nil, errors.New("seriesId must be provided")

	}
	url += "seriesId=" + *seriesId

	if seasonNumber != nil && *seasonNumber != "" {
		url += "&seasonNumber=" + *seasonNumber
	}

	fmt.Println(url)

	body, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	episodes := []models.Episode{}
	err = json.Unmarshal(body, &episodes)

	if err != nil {
		return nil, err
	}

	return episodes, nil
}

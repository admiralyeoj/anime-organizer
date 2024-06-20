package sonarr

import (
	"encoding/json"
	"fmt"

	"github.com/admiralyeoj/anime-organizer/internal/sonarr/models"
)

func (c *Client) GetSeries() ([]models.Series, error) {
	url := c.baseUrl + "/v3/series"
	fmt.Println(url)

	body, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	seriesResp := []models.Series{}
	err = json.Unmarshal(body, &seriesResp)
	if err != nil {
		return nil, err
	}

	return seriesResp, nil
}

package sonarr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/admiralyeoj/anime-organizer/internal/sonarr/models"
)

func (c *Client) GetSeries() ([]models.Series, error) {
	url := os.Getenv("SONARR_URL")
	method := "GET"

	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("X-Api-Key", os.Getenv("SONARR_KEY"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	seriesResp := []models.Series{}
	err = json.Unmarshal(body, &seriesResp)

	if err != nil {
		return nil, err
	}

	return seriesResp, nil
}

package sonarr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RespSeries struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	SortTitle  *string    `json:"sortTitle"`
	Status     string     `json:"status"`
	Overview   string     `json:"overview"`
	Images     []Image    `json:"images"`
	Seasons    []Season   `json:"seasons"`
	Year       int16      `json:"year"`
	TvdbId     int        `json:"tvdbId"`
	ImdbId     string     `json:"imdbId"`
	TitleSlug  string     `json:"titleSlug"`
	Genres     []string   `json:"genres"`
	Statistics Statistics `json:"statistics"`
}

type Image struct {
	CoverType string `json:"seasonNumber"`
	Url       string `json:"remoteUrl"`
}

type Season struct {
	SeasonNumber int8       `json:"seasonNumber"`
	Statistics   Statistics `json:"statistics"`
}

type Statistics struct {
	SeasonCount       int `json:"seasonCount"`
	EpisodeCount      int `json:"episodeCount"`
	TotalEpisodeCount int `json:"totalEpisodeCount"`
	PercentOfEpisodes int `json:"percentOfEpisodes"`
}

func (c *Client) GetSeries() ([]byte, error) {
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

	seriesResp := []RespSeries{}
	err = json.Unmarshal(body, &seriesResp)

	if err != nil {
		return nil, err
	}

	// Iterate over each series in seriesResp
	for _, series := range seriesResp {
		fmt.Println("Series Information:")
		fmt.Println("ID:", series.Id)
		fmt.Println("Title:", series.Title)
		if series.SortTitle != nil {
			fmt.Println("Sort Title:", *series.SortTitle)
		}
		fmt.Println("Status:", series.Status)
		fmt.Println("Overview:", series.Overview)
		fmt.Println("Year:", series.Year)
		fmt.Println("TVDB ID:", series.TvdbId)
		fmt.Println("IMDB ID:", series.ImdbId)
		fmt.Println("Title Slug:", series.TitleSlug)

		fmt.Println("\nGenres:")
		for _, genre := range series.Genres {
			fmt.Println("-", genre)
		}

		fmt.Println("\nSeasons:")
		for _, season := range series.Seasons {
			fmt.Println("Season Number:", season.SeasonNumber)
			fmt.Println("  Season Count:", season.Statistics.SeasonCount)
			fmt.Println("  Episode Count:", season.Statistics.EpisodeCount)
			fmt.Println("  Total Episode Count:", season.Statistics.TotalEpisodeCount)
			fmt.Println("  Percent of Episodes:", season.Statistics.PercentOfEpisodes)
		}

		fmt.Println("\nOverall Statistics:")
		fmt.Println("Season Count:", series.Statistics.SeasonCount)
		fmt.Println("Episode Count:", series.Statistics.EpisodeCount)
		fmt.Println("Total Episode Count:", series.Statistics.TotalEpisodeCount)
		fmt.Println("Percent of Episodes:", series.Statistics.PercentOfEpisodes)

		fmt.Println()
	}

	// for _, value := range body {
	// 	fmt.Println(value)
	// 	break
	// }

	return body, nil
}

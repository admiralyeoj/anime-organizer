package models

type Series struct {
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

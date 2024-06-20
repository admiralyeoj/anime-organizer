package models

type Season struct {
	SeasonNumber int8       `json:"seasonNumber"`
	Statistics   Statistics `json:"statistics"`
}

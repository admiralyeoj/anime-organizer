package models

type Statistics struct {
	SeasonCount       int `json:"seasonCount"`
	EpisodeCount      int `json:"episodeCount"`
	TotalEpisodeCount int `json:"totalEpisodeCount"`
	PercentOfEpisodes int `json:"percentOfEpisodes"`
}

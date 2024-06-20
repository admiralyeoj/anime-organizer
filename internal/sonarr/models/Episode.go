package models

type Episode struct {
	Id                    int    `json:"id"`
	SeriesId              int    `json:"seriesId"`
	TvdbId                int    `json:"tvdbId"`
	SeasonNumber          int    `json:"seasonNumber"`
	EpisodeNumber         int    `json:"episodeNumber"`
	Title                 string `json:"title"`
	AirDate               string `json:"airDate"`
	Overview              string `json:"overview"`
	AbsoluteEpisodeNumber int    `json:"absoluteEpisodeNumber"`
	SceneEpisodeNumber    int    `json:"sceneEpisodeNumber"`
	SceneSeasonNumber     int    `json:"sceneSeasonNumber"`
}

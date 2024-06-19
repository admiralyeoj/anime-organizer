package sonarr

type Config struct {
	SonarrApiClient Client
	host            string
	apiKey          string
}

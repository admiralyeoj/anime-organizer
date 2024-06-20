package sonarr

import (
	"net/http"
	"os"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	baseUrl    string
	apiKey     string
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseUrl: os.Getenv("SONARR_URL"),
		apiKey:  os.Getenv("SONARR_KEY"),
	}
}

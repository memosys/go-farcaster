package neynar

import (
	v1 "go-farcaster/pkg/v1"
	v2 "go-farcaster/pkg/v2"
	"net/http"
)

type APIClient struct {
	logger *Logger
	v1     *v1.NeynarV1APIClient
	v2     *v2.NeynarV2APIClient
}

func NewNeynarAPIClient(apiKey string, logger *Logger, client *http.Client) *APIClient {
	return &APIClient{
		logger: logger,
		v1:     v1.NewV1Client(apiKey, client),
		v2:     v2.NewV2Client(apiKey, client),
	}
}

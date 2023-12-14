package neynar

import (
	"go-farcaster/pkg/v1"
	"go-farcaster/pkg/v2"
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
		v1:     v1.NewV1Client(apiKey, logger, client),
		v2:     v2.NewV2Client(apiKey, logger, client),
	}
}

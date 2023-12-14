package v2

import (
	logger "go-farcaster/pkg/neynar"
	"net/http"
)

// API interfaces for v2 endpoints
type SignerApi interface{}
type FeedApi interface{}
type CastApi interface{}
type UserApi interface{}
type ReactionApi interface{}
type FollowsApi interface{}

// NeynarV2APIClient struct with APIs and Logger.
type NeynarV2APIClient struct {
	logger *logger.Logger
	apis   struct {
		signer   SignerApi
		feed     FeedApi
		cast     CastApi
		user     UserApi
		reaction ReactionApi
		follow   FollowsApi
	}
}

func NewV2Client(apiKey string, logger *logger.Logger, client *http.Client) *NeynarV2APIClient {
	// Initialize API clients here with apiKey and http.Client
	return &NeynarV2APIClient{
		logger: logger,
		// Initialize APIs
	}
}

// Method implementations for NeynarV2APIClient such as fetchSigner, createSigner, fetchFeed, etc.

package v2

import (
	"net/http"
)

// API interfaces for v2 endpoints
type SignerApi interface{}
type FeedApi interface{}
type CastApi interface{}
type UserApi interface{}
type ReactionApi interface{}
type FollowsApi interface{}

// NeynarV2APIClient struct with APIs and neynar.
type NeynarV2APIClient struct {
	apis struct {
		signer   SignerApi
		feed     FeedApi
		cast     CastApi
		user     UserApi
		reaction ReactionApi
		follow   FollowsApi
	}
}

func NewV2Client(apiKey string, client *http.Client) *NeynarV2APIClient {
	// Initialize API clients here with apiKey and http.Client
	return &NeynarV2APIClient{

		// Initialize APIs
	}
}

// Method implementations for NeynarV2APIClient such as fetchSigner, createSigner, fetchFeed, etc.

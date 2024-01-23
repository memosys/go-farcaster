package v2

import (
	v2 "go-farcaster/pkg/openapi/openapiv2"
	"net/http"
)

// NeynarV2APIClient struct with APIs and neynar.
type NeynarV2APIClient struct {
	apis struct {
		signer        v2.SignerAPIService
		feed          v2.FeedAPIService
		cast          v2.CastAPIService
		user          v2.UserAPIService
		reaction      v2.ReactionAPIService
		follow        v2.FollowsAPIService
		storage       v2.StorageAPIService
		notifications v2.NotificationsAPIService
	}
}

func NewV2Client(apiKey string, client *http.Client) *NeynarV2APIClient {
	// Initialize API clients here with apiKey and http.Client
	return &NeynarV2APIClient{

		// Initialize APIs
	}
}

// Method implementations for NeynarV2APIClient such as fetchSigner, createSigner, fetchFeed, etc.

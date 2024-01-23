package v1

import (
	v1 "go-farcaster/pkg/openapi/openapiv1"
	"net/http"
)

type NeynarV1APIClient struct {
	apis struct {
		user          *v1.UserAPIService
		cast          *v1.CastAPIService
		verification  *v1.VerificationAPIService
		notifications *v1.NotificationsAPIService
		reactions     *v1.ReactionsAPIService
		follows       *v1.FollowsAPIService
	}
}

// NewNeynarV1APIClient function to instantiate NeynarV1APIClient.
func NewV1Client(apiKey string, client *http.Client) *NeynarV1APIClient {
	// Initialize API clients here with apiKey and http.Client
	return &NeynarV1APIClient{

		// Initialize APIs
	}
}

// Additional methods for NeynarV1APIClient should be implemented here.
// For example: fetchCastsInThread, fetchCastsForUser, etc.

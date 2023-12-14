package v1

import (
	"net/http"
)

// API interfaces for v1 endpoints
type UserApi interface{}
type CastApi interface{}
type VerificationApi interface{}
type NotificationsApi interface{}
type ReactionsApi interface{}
type FollowsApi interface{}

type NeynarV1APIClient struct {
	apis struct {
		user          UserApi
		cast          CastApi
		verification  VerificationApi
		notifications NotificationsApi
		reactions     ReactionsApi
		follows       FollowsApi
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

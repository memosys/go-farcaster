/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// NotificationsAPIService NotificationsAPI service
type NotificationsAPIService service

type ApiMentionsAndRepliesRequest struct {
	ctx context.Context
	ApiService *NotificationsAPIService
	apiKey *string
	fid *int32
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiMentionsAndRepliesRequest) ApiKey(apiKey string) ApiMentionsAndRepliesRequest {
	r.apiKey = &apiKey
	return r
}

// fid of a user
func (r ApiMentionsAndRepliesRequest) Fid(fid int32) ApiMentionsAndRepliesRequest {
	r.fid = &fid
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiMentionsAndRepliesRequest) ViewerFid(viewerFid int32) ApiMentionsAndRepliesRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 150)
func (r ApiMentionsAndRepliesRequest) Limit(limit int32) ApiMentionsAndRepliesRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiMentionsAndRepliesRequest) Cursor(cursor string) ApiMentionsAndRepliesRequest {
	r.cursor = &cursor
	return r
}

func (r ApiMentionsAndRepliesRequest) Execute() (*MentionsAndRepliesResponse, *http.Response, error) {
	return r.ApiService.MentionsAndRepliesExecute(r)
}

/*
MentionsAndReplies Get mentions and replies

Gets a list of mentions and replies to the user’s casts in reverse chronological order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMentionsAndRepliesRequest
*/
func (a *NotificationsAPIService) MentionsAndReplies(ctx context.Context) ApiMentionsAndRepliesRequest {
	return ApiMentionsAndRepliesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MentionsAndRepliesResponse
func (a *NotificationsAPIService) MentionsAndRepliesExecute(r ApiMentionsAndRepliesRequest) (*MentionsAndRepliesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MentionsAndRepliesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.MentionsAndReplies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/mentions-and-replies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 25
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReactionsAndRecastsRequest struct {
	ctx context.Context
	ApiService *NotificationsAPIService
	apiKey *string
	fid *int32
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiReactionsAndRecastsRequest) ApiKey(apiKey string) ApiReactionsAndRecastsRequest {
	r.apiKey = &apiKey
	return r
}

// fid of a user
func (r ApiReactionsAndRecastsRequest) Fid(fid int32) ApiReactionsAndRecastsRequest {
	r.fid = &fid
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiReactionsAndRecastsRequest) ViewerFid(viewerFid int32) ApiReactionsAndRecastsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 150)
func (r ApiReactionsAndRecastsRequest) Limit(limit int32) ApiReactionsAndRecastsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiReactionsAndRecastsRequest) Cursor(cursor string) ApiReactionsAndRecastsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiReactionsAndRecastsRequest) Execute() (*ReactionsAndRecastsResponse, *http.Response, error) {
	return r.ApiService.ReactionsAndRecastsExecute(r)
}

/*
ReactionsAndRecasts Get reactions and recasts

Get a list of reactions and recasts to the users’s casts in reverse chronological order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReactionsAndRecastsRequest
*/
func (a *NotificationsAPIService) ReactionsAndRecasts(ctx context.Context) ApiReactionsAndRecastsRequest {
	return ApiReactionsAndRecastsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReactionsAndRecastsResponse
func (a *NotificationsAPIService) ReactionsAndRecastsExecute(r ApiReactionsAndRecastsRequest) (*ReactionsAndRecastsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReactionsAndRecastsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.ReactionsAndRecasts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/reactions-and-recasts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 25
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

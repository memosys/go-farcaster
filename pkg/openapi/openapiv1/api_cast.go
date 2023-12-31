/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CastAPIService CastAPI service
type CastAPIService service

type ApiAllCastsInThreadRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	threadHash *string
	viewerFid *int32
}

// API key required for authentication.
func (r ApiAllCastsInThreadRequest) ApiKey(apiKey string) ApiAllCastsInThreadRequest {
	r.apiKey = &apiKey
	return r
}

// The hash of the thread to retrieve casts from.
func (r ApiAllCastsInThreadRequest) ThreadHash(threadHash string) ApiAllCastsInThreadRequest {
	r.threadHash = &threadHash
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiAllCastsInThreadRequest) ViewerFid(viewerFid int32) ApiAllCastsInThreadRequest {
	r.viewerFid = &viewerFid
	return r
}

func (r ApiAllCastsInThreadRequest) Execute() (*AllCastsInThreadResponse, *http.Response, error) {
	return r.ApiService.AllCastsInThreadExecute(r)
}

/*
AllCastsInThread Retrieve all casts in a given thread hash

Gets all casts, including root cast and all replies for a given thread hash. No limit the depth of replies.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAllCastsInThreadRequest
*/
func (a *CastAPIService) AllCastsInThread(ctx context.Context) ApiAllCastsInThreadRequest {
	return ApiAllCastsInThreadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AllCastsInThreadResponse
func (a *CastAPIService) AllCastsInThreadExecute(r ApiAllCastsInThreadRequest) (*AllCastsInThreadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllCastsInThreadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.AllCastsInThread")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/all-casts-in-thread"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.threadHash == nil {
		return localVarReturnValue, nil, reportError("threadHash is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "threadHash", r.threadHash, "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
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

type ApiCastRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	hash *string
	viewerFid *int32
}

// API key required for authentication.
func (r ApiCastRequest) ApiKey(apiKey string) ApiCastRequest {
	r.apiKey = &apiKey
	return r
}

// Cast hash
func (r ApiCastRequest) Hash(hash string) ApiCastRequest {
	r.hash = &hash
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiCastRequest) ViewerFid(viewerFid int32) ApiCastRequest {
	r.viewerFid = &viewerFid
	return r
}

func (r ApiCastRequest) Execute() (*CastResponse, *http.Response, error) {
	return r.ApiService.CastExecute(r)
}

/*
Cast Retrieve cast for a given hash

Gets information about an individual cast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastRequest
*/
func (a *CastAPIService) Cast(ctx context.Context) ApiCastRequest {
	return ApiCastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastResponse
func (a *CastAPIService) CastExecute(r ApiCastRequest) (*CastResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.Cast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.hash == nil {
		return localVarReturnValue, nil, reportError("hash is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
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

type ApiCastsRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	fid *int32
	parentUrl *string
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiCastsRequest) ApiKey(apiKey string) ApiCastsRequest {
	r.apiKey = &apiKey
	return r
}

// fid of a user
func (r ApiCastsRequest) Fid(fid int32) ApiCastsRequest {
	r.fid = &fid
	return r
}

// A cast can be part of a certain channel. The channel is identified by &#x60;parent_url&#x60;. All casts in the channel ladder up to the same parent_url.
func (r ApiCastsRequest) ParentUrl(parentUrl string) ApiCastsRequest {
	r.parentUrl = &parentUrl
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiCastsRequest) ViewerFid(viewerFid int32) ApiCastsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 150)
func (r ApiCastsRequest) Limit(limit int32) ApiCastsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiCastsRequest) Cursor(cursor string) ApiCastsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiCastsRequest) Execute() (*CastsResponse, *http.Response, error) {
	return r.ApiService.CastsExecute(r)
}

/*
Casts Retrieve casts for a given user

Gets the most recent casts for a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastsRequest
*/
func (a *CastAPIService) Casts(ctx context.Context) ApiCastsRequest {
	return ApiCastsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastsResponse
func (a *CastAPIService) CastsExecute(r ApiCastsRequest) (*CastsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.Casts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/casts"

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
	if r.parentUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parent_url", r.parentUrl, "")
	}
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

type ApiRecentCastsRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiRecentCastsRequest) ApiKey(apiKey string) ApiRecentCastsRequest {
	r.apiKey = &apiKey
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiRecentCastsRequest) ViewerFid(viewerFid int32) ApiRecentCastsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 100)
func (r ApiRecentCastsRequest) Limit(limit int32) ApiRecentCastsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiRecentCastsRequest) Cursor(cursor string) ApiRecentCastsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiRecentCastsRequest) Execute() (*RecentCastsResponse, *http.Response, error) {
	return r.ApiService.RecentCastsExecute(r)
}

/*
RecentCasts Get Recent Casts

Get a list of casts from the protocol in reverse chronological order based on timestamp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRecentCastsRequest
*/
func (a *CastAPIService) RecentCasts(ctx context.Context) ApiRecentCastsRequest {
	return ApiRecentCastsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RecentCastsResponse
func (a *CastAPIService) RecentCastsExecute(r ApiRecentCastsRequest) (*RecentCastsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RecentCastsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.RecentCasts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/recent-casts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}

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

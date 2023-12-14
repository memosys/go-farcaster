# \NotificationsAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Notifications**](NotificationsAPI.md#Notifications) | **Get** /farcaster/notifications | Retrieve notifications for a given user
[**NotificationsChannel**](NotificationsAPI.md#NotificationsChannel) | **Get** /farcaster/notifications/channel | Retrieve notifications for a user for a given channel



## Notifications

> NotificationsResponse Notifications(ctx).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Retrieve notifications for a given user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
    fid := int32(56) // int32 | 
    limit := int32(56) // int32 | Number of results to retrieve (default 25, max 50) (optional) (default to 25)
    cursor := "cursor_example" // string | Pagination cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.Notifications(context.Background()).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.Notifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Notifications`: NotificationsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.Notifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **fid** | **int32** |  | 
 **limit** | **int32** | Number of results to retrieve (default 25, max 50) | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsChannel

> NotificationsResponse NotificationsChannel(ctx).ApiKey(apiKey).Fid(fid).ParentUrls(parentUrls).Limit(limit).Cursor(cursor).Execute()

Retrieve notifications for a user for a given channel



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
    fid := int32(56) // int32 | 
    parentUrls := "chain://eip155:1/erc721:0xd4498134211baad5846ce70ce04e7c4da78931cc" // string | Comma separated channel parent_urls (find mappings here - https://github.com/neynarxyz/farcaster-channels/blob/main/warpcast.json)
    limit := int32(56) // int32 | Number of results to retrieve (default 25, max 50) (optional) (default to 25)
    cursor := "cursor_example" // string | Pagination cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.NotificationsChannel(context.Background()).ApiKey(apiKey).Fid(fid).ParentUrls(parentUrls).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.NotificationsChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsChannel`: NotificationsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.NotificationsChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **fid** | **int32** |  | 
 **parentUrls** | **string** | Comma separated channel parent_urls (find mappings here - https://github.com/neynarxyz/farcaster-channels/blob/main/warpcast.json) | 
 **limit** | **int32** | Number of results to retrieve (default 25, max 50) | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


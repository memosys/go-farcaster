# \ReactionAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReaction**](ReactionAPI.md#DeleteReaction) | **Delete** /farcaster/reaction | Delete a reaction
[**PostReaction**](ReactionAPI.md#PostReaction) | **Post** /farcaster/reaction | Posts a reaction
[**ReactionsUser**](ReactionAPI.md#ReactionsUser) | **Get** /farcaster/reactions/user | Fetches reactions for a given user



## DeleteReaction

> OperationResponse DeleteReaction(ctx).ApiKey(apiKey).ReactionReqBody(reactionReqBody).Execute()

Delete a reaction



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
    reactionReqBody := *openapiclient.NewReactionReqBody("SignerUuid_example", openapiclient.ReactionType("like"), "Target_example") // ReactionReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionAPI.DeleteReaction(context.Background()).ApiKey(apiKey).ReactionReqBody(reactionReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.DeleteReaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReaction`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.DeleteReaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **reactionReqBody** | [**ReactionReqBody**](ReactionReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReaction

> OperationResponse PostReaction(ctx).ApiKey(apiKey).ReactionReqBody(reactionReqBody).Execute()

Posts a reaction



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
    reactionReqBody := *openapiclient.NewReactionReqBody("SignerUuid_example", openapiclient.ReactionType("like"), "Target_example") // ReactionReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionAPI.PostReaction(context.Background()).ApiKey(apiKey).ReactionReqBody(reactionReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.PostReaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostReaction`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.PostReaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **reactionReqBody** | [**ReactionReqBody**](ReactionReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsUser

> ReactionsResponse ReactionsUser(ctx).ApiKey(apiKey).Fid(fid).Type_(type_).Limit(limit).Cursor(cursor).Execute()

Fetches reactions for a given user



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
    fid := int32(3) // int32 | 
    type_ := openapiclient.ReactionsType("likes") // ReactionsType | Type of reaction to fetch (likes or recasts or all)
    limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
    cursor := "cursor_example" // string | Pagination cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionAPI.ReactionsUser(context.Background()).ApiKey(apiKey).Fid(fid).Type_(type_).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.ReactionsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsUser`: ReactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.ReactionsUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReactionsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **fid** | **int32** |  | 
 **type_** | [**ReactionsType**](ReactionsType.md) | Type of reaction to fetch (likes or recasts or all) | 
 **limit** | **int32** | Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ReactionsResponse**](ReactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

